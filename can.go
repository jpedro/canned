package canned

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/jpedro/crypto"
	// "gopkg.in/yaml.v2"
)

// Can struct
type Can struct {
	file     string
	password string
	Version  string          `json:"version" yaml:"version"`
	Metadata Metadata        `json:"metadata" yaml:"metadata"`
	Items    map[string]Item `json:"items" yaml:"items"`
}

// Loads a can file into memory
func (can *Can) load() error {
	content, err := ioutil.ReadFile(can.file)
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

	return nil
}

// Saves a can file
func (can *Can) Save() error {
	// data, err := json.MarshalIndent(can, "", "  ")
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

	err = ioutil.WriteFile(can.file, []byte(headed), 0644)
	if err != nil {
		return err
	}

	// dataY, err := yaml.Marshal(can)
	// if err != nil {
	//     return err
	// }
	// err = ioutil.WriteFile(can.File + ".json", data, 0644)
	// if err != nil {
	//     return err
	// }
	// err = ioutil.WriteFile(can.File + ".yaml", dataY, 0644)
	// if err != nil {
	//     return err
	// }
	// var loaded *Can
	// err = json.Unmarshal(data, &loaded)
	// if err != nil {
	//     return err
	// }
	// var loadedY *Can
	// err = yaml.Unmarshal(dataY, &loadedY)
	// if err != nil {
	//     return err
	// }

	return nil
}

// Stores an item
func (can *Can) SetItem(name string, value string) error {
	var item *Item
	item, err := NewItem(name, value)
	if err != nil {
		return err
	}

	can.Items[name] = *item

	return nil
}

// Renames an existing item
func (can *Can) RenameItem(name string, new string) error {
	item, exists := can.Items[name]
	if !exists {
		return fmt.Errorf("Item %s doesn't exist.", name)
	}

	// item.Name = new
	can.Items[new] = item
	delete(can.Items, name)

	return nil
}

// Gets an existing item
func (can *Can) GetItem(name string) (*Item, error) {
	item, exists := can.Items[name]
	if !exists {
		return nil, fmt.Errorf("Item %s doesn't exist.", name)
	}

	return &item, nil
}

func (can *Can) DelItem(name string) error {
	_, exists := can.Items[name]
	if !exists {
		return fmt.Errorf("Item %s doesn't exist.", name)
	}

	delete(can.Items, name)

	return nil
}

// Appends a tag to an item
func (can *Can) AddTag(name string, tag string) error {
	item, err := can.GetItem(name)
	if err != nil {
		return fmt.Errorf("Item %s doesn't exist.", name)
	}

	if exists(item.Tags, tag) {
		return nil
	}

	item.Metadata.UpdatedAt = time.Now()
	item.Tags = append(item.Tags, tag)
	can.Items[name] = *item

	return nil
}

// Removes a tag from an item
func (can *Can) DelTag(name string, tag string) bool {
	item, err := can.GetItem(name)
	if err != nil {
		return false
	}

	if !exists(item.Tags, tag) {
		return false
	}

	item.Metadata.UpdatedAt = time.Now()
	item.Tags = remove(item.Tags, tag)
	can.Items[name] = *item

	return true
}
