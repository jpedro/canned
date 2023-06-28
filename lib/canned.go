package canned

import (
	"fmt"
	"time"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	VERSION   = "v1"
	ALGORITHM = "GCM"
	SEPARATOR = "\n\n"
)

var (
	db *sql.DB
)

func openDatabase(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	return db, err
}

func initDatabase(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic(err)
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
			("created", CURRENT_TIMESTAMP)
		`, VERSION, ALGORITHM),

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
			value TEXT NOT NULL,
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

	return db, nil
}

// Creates a new can object
func NewCan(file string, password string) (*Can, error) {
	err := ensurePasswordExists(password)
	if err != nil {
		return nil, err
	}

	can := &Can{}
	can.file = file
	can.password = password
	can.Version = VERSION
	can.Algorithm = ALGORITHM
	can.Metadata.CreatedAt = time.Now()
	can.Items = make(map[string]*Item)

	return can, nil
}

// Creates a new item object
func NewItem(name string, content string) (*Item, error) {
	if name == "" {
		return nil, fmt.Errorf("name can't be empty")
	}

	item := &Item{}
	item.Content = content
	item.Metadata.CreatedAt = time.Now()
	// item.Tags = []string{}

	return item, nil
}

// Opens a can file
// If the file doesn't exist, it will fail to load
func OpenCan(file string, password string) (*Can, error) {
	can, err := NewCan(file, password)
	if err != nil {
		return nil, err
	}
	err = can.load()
	if err != nil {
		return nil, err
	}

	db, err = openDatabase(file + ".sqlite3")
	if err != nil {
		panic(err)
	}

	return can, err
}

// Initializes a new can file
// This will overwrite any existing file.
func InitCan(file string, password string) (*Can, error) {
	can, err := NewCan(file, password)
	if err != nil {
		return nil, err
	}

	db, err = initDatabase(file + ".sqlite")
	if err != nil {
		return nil, err
	}

	err = can.Save()
	if err != nil {
		return nil, err
	}

	return can, err
}
