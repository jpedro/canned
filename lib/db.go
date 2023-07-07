package canned

import (
	"fmt"
	"strings"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func openDatabase(file, password string) error {
	var err error
	db, err = sql.Open("sqlite3", file)
	if err != nil {
		return err
	}
	verifyDatabase()
	return err
}

func verifyDatabase() error {
	return nil
}

func initDatabase(file, password string) error {
	var err error
	db, err = sql.Open("sqlite3", file)
	if err != nil {
		return err
	}

	verification, err := encrypt(password, password)
	if err != nil {
		return err
	}

	queries := []string{
		`PRAGMA foreign_keys = ON`,

		`DROP TABLE IF EXISTS header`,
		`CREATE TABLE header (
			name TEXT PRIMARY KEY NOT NULL,
			value TEXT NOT NULL
		)`,
		fmt.Sprintf(`INSERT INTO header (name, value) VALUES
			("version", "%s"),
			("algorithm", "%s"),
			("verification", "%s"),
			("created", CURRENT_TIMESTAMP)
		`, VERSION, ALGORITHM, verification),

		`DROP TABLE IF EXISTS item`,
		`CREATE TABLE item (
			id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			type TEXT NOT NULL,
			ref INTEGER DEFAULT 0,
			name TEXT NOT NULL UNIQUE,
			value TEXT NOT NULL,
			strength INTEGER NOT NULL,
			created DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated DATETIME,
			tags TEXT
		)`,

		`DROP TABLE IF EXISTS file`,
		`CREATE TABLE file (
			id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			type TEXT NOT NULL,
			name TEXT NOT NULL UNIQUE,
			value BLOB NOT NULL,
			size INTEGER NOT NULL,
			created DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated DATETIME
		)`,

		`DROP TABLE IF EXISTS tag`,
		`CREATE TABLE tag (
			id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			name TEXT NOT NULL UNIQUE,
			created DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated DATETIME
		)`,

		`DROP TABLE IF EXISTS item_tag`,
		`CREATE TABLE item_tag (
			item_id TEXT NOT NULL
				REFERENCES items(id)
				ON DELETE CASCADE,
			tag_id TEXT NOT NULL
				REFERENCES tags(id)
				ON DELETE CASCADE,
			created DATETIME NOT NULL,
			PRIMARY KEY(item_id, tag_id)
		)`,
	}

	for _, query := range queries {
		res, err2 := db.Exec(query)
		if err2 != nil {
			panic(err2)
		}
		affected, err3 := res.RowsAffected()
		if err3 != nil {
			panic(err3)
		}
		fmt.Println("affected:", query, affected)
	}

	return nil
}

func saveDatabase(can *Can) error {
	err := openDatabase(can.file+DB_EXT, can.password)
	if err != nil {
		return err
	}

	if len(can.Items) < 1 {
		return nil
	}

	params := []any{}
	values := []string{}
	for name, item := range can.Items {
		holders := "(?, ?, ?, ?, ?, ?, ?)"
		values = append(values, holders)
		encrypted, err := encrypt(item.Content, can.password)
		if err != nil {
			panic(err)
		}

		params = append(params, []any{
			name,
			"secret",
			encrypted,
			strength(item.Content),
			item.Metadata.Created,
			item.Metadata.Updated,
			strings.Join(item.Tags, ", "),
		})
	}

	query := `
		INSERT INTO
				header
				(name, value)
		VALUES
				('updated', CURRENT_TIMESTAMP)
		ON CONFLICT(name)
			DO UPDATE SET
			value = excluded.value
	`
	execQuery(query)

	query = `
	INSERT INTO
			item
			(name, type, value, strength, created, updated, tags)
	VALUES
	` + strings.Join(values, ", ")

	query = query + `
		ON CONFLICT(name)
		DO UPDATE SET
			value = excluded.value,
			updated = CURRENT_TIMESTAMP,
			tags = excluded.tags
	`

	fmt.Println("params:", params)
	params = Flatten(params)
	fmt.Println("params:", params)
	execQuery(query, params...)

	return err
}

func execQuery(query string, params ...any) {
	fmt.Println("query:", query)
	fmt.Println("params:", params)

	res, err := db.Exec(query, params...)
	if err != nil {
		panic(err)
	}

	affected, err2 := res.RowsAffected()
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("affected:", query, affected)
}
