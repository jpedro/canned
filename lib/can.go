package canned

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

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

	err = can.dump(data)

	return err
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
