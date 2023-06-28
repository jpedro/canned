package canned

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"database/sql"
	"github.com/jpedro/crypto"
	"gopkg.in/yaml.v2"
)

// Can struct
type Can struct {
	file      string
	password  string
	Version   string           `json:"version"   yaml:"version"`
	Algorithm string           `json:"algorithm" yaml:"algorithm"`
	Metadata  Metadata         `json:"metadata"  yaml:"metadata"`
	Items     map[string]*Item `json:"items"     yaml:"items"`
}

// Loads the can file into memory
func (can *Can) load() error {
	content, err := os.ReadFile(can.file)
	if err != nil {
		return err
	}

	headers, payload := getHeaders(string(content))
	striped := strip(payload)
	decrypted, err := crypto.Decrypt(striped, can.password)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(decrypted), &can)
	if err != nil {
		return err
	}

	can.Version = headers["version"]
	can.Algorithm = headers["version"]

	return nil
}

// Saves the can into the file
func (can *Can) Save() error {
	data, err := json.Marshal(can)
	if err != nil {
		return err
	}

	encrypted, err := crypto.Encrypt(string(data), can.password)
	if err != nil {
		return err
	}

	chunks := align(encrypted, 64)
	aligned := strings.Join(chunks, "\n")
	headed := addHeaders(aligned)

	err = os.WriteFile(can.file, []byte(headed), 0644)
	if err != nil {
		return err
	}

	err = can.saveDatabase()
	if err != nil {
		return err
	}

	err = can.dump(data)
	if err != nil {
		return err
	}

	return err
}

func Flatten(nested []any) []any {
	flattened := make([]any, 0)

	for _, i := range nested {
		switch i.(type) {
		case []interface{}:
			flattenedSubArray := Flatten(i.([]any))
			flattened = append(flattened, flattenedSubArray...)
		case interface{}:
			flattened = append(flattened, i)
		}
	}

	return flattened
}

func (can *Can) saveDatabase() error {
	db, err := openDatabase(can.file + ".sqlite")
	if err != nil {
		panic(err)
	}

	if len(can.Items) < 1 {
		return nil
	}

	params := []any{}
	values := []string{}
	for name, item := range can.Items {
		holders := "(?, ?, ?, ?, ?, ?, ?)"
		values = append(values, holders)
		encrypted, err := can.encrypt(item.Content)
		if err != nil {
			panic(err)
		}

		params = append(params, []any{
			name,
			"secret",
			encrypted,
			can.strength(item.Content),
			item.Metadata.CreatedAt,
			item.Metadata.UpdatedAt,
			strings.Join(item.Tags, ", "),
		})
	}

	query := `
		INSERT INTO
				header (name, value)
		VALUES
				('updated', CURRENT_TIMESTAMP)
		ON CONFLICT(name)
			DO UPDATE SET
			value = excluded.value
	`
	can.execQuery(db, query)

	query = `
	INSERT INTO
			item (name, type, value, strength, created, updated, tags)
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
	can.execQuery(db, query, params...)

	return err
}

func (can *Can) execQuery(db *sql.DB, query string, params ...any) {
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

func (can *Can) dump(data []byte) error {
	dump := env("CANNED_DUMP", "")
	if dump != "yes-pretty-please-dump-the-can" {
		return nil
	}

	redacted := can
	for name := range can.Items {
		redacted.SetItem(name, "[redacted]")
	}

	dataJson, err := json.Marshal(redacted)
	if err != nil {
		return err
	}

	dataYaml, err := yaml.Marshal(redacted)
	if err != nil {
		return err
	}

	err = os.WriteFile(can.file+".json", dataJson, 0644)
	if err != nil {
		return err
	}

	err = os.WriteFile(can.file+".yaml", dataYaml, 0644)
	if err != nil {
		return err
	}

	var loadedJson *Can
	err = json.Unmarshal(data, &loadedJson)
	if err != nil {
		return err
	}

	var loadedYaml *Can
	err = yaml.Unmarshal(dataYaml, &loadedYaml)
	if err != nil {
		return err
	}

	return nil
}

// SetItem stores an item's name and value
func (can *Can) SetItem(name string, value string) error {
	var item *Item

	item = can.Items[name]
	if item != nil {
		item.Content = value
		item.Metadata.UpdatedAt = time.Now()
		can.Metadata.UpdatedAt = item.Metadata.UpdatedAt
		return nil
	}

	item, err := NewItem(name, value)
	if err != nil {
		return err
	}

	can.Items[name] = item

	return nil
}

// RenameItem renames an existing item
func (can *Can) RenameItem(name string, newName string) error {
	item, exists := can.Items[name]
	if !exists {
		return fmt.Errorf("Item %s doesn't exist", name)
	}

	newItem := item
	newItem.Metadata.UpdatedAt = time.Now()
	can.Metadata.UpdatedAt = newItem.Metadata.UpdatedAt

	can.Items[newName] = newItem
	delete(can.Items, name)

	return nil
}

// GetItem retrieves an existing item
func (can *Can) GetItem(name string) (*Item, error) {
	item, exists := can.Items[name]
	if !exists {
		return nil, fmt.Errorf("Item %s doesn't exist", name)
	}

	return item, nil
}

// DelItem deletes an existing item
func (can *Can) DelItem(name string) error {
	_, exists := can.Items[name]
	if !exists {
		return fmt.Errorf("Item %s doesn't exist", name)
	}

	can.Metadata.UpdatedAt = time.Now()
	delete(can.Items, name)

	return nil
}

// AddTag appends a tag to an item
func (can *Can) AddTag(name string, tag string) error {
	item, err := can.GetItem(name)
	if err != nil {
		return fmt.Errorf("Item %s doesn't exist", name)
	}

	if exists(item.Tags, tag) {
		return nil
	}

	item.Metadata.UpdatedAt = time.Now()
	can.Metadata.UpdatedAt = item.Metadata.UpdatedAt
	item.Tags = append(item.Tags, tag)
	can.Items[name] = item

	return nil
}

// DelTag removes a tag from an item
func (can *Can) DelTag(name string, tag string) error {
	item, err := can.GetItem(name)
	if err != nil {
		return err
	}

	if !exists(item.Tags, tag) {
		return fmt.Errorf("Item %s tag %s doesn't exist", name, tag)
	}

	item.Tags = remove(item.Tags, tag)
	item.Metadata.UpdatedAt = time.Now()
	can.Metadata.UpdatedAt = item.Metadata.UpdatedAt
	can.Items[name] = item

	return nil
}

func (can *Can) encrypt(data string) (string, error) {
	encrypted, err := crypto.Encrypt(string(data), can.password)
	if err != nil {
		return "", err
	}

	return encrypted, nil
}

func (can *Can) decrypt(data string) (string, error) {
	decrypted, err := crypto.Decrypt(string(data), can.password)
	if err != nil {
		return "", err
	}

	return decrypted, nil
}

func (can *Can) strength(value string) int {
	mod := len(value) / 10
	if mod < 1 {
		mod = 1
	}

	return mod
}
