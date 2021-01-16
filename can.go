package can

import (
    "fmt"
    "time"
    // "errors"
    "strings"
    "encoding/json"
    "io/ioutil"

    "github.com/jpedro/crypto"
)

type Store struct {
    File        string              `json:"-"`
    Version     string              `json:"version"`
    Metadata    Metadata            `json:"metadata"`
    Items       map[string]Item     `json:"items"`
}

func (store *Store) Load() error {
    data, err := ioutil.ReadFile(store.File + ".json")
    if err != nil {
        return err
    }

    err = json.Unmarshal(data, &store)
    if err != nil {
        return err
    }

    return nil
}

func (store *Store) Save() error {
    data, err := json.MarshalIndent(store, "", "  ")
    if err != nil {
        return err
    }

    encrypted, err := crypto.Encrypt(string(data), "pp")
    if err != nil {
        return err
    }

    chunks    := align(encrypted, 64)
    aligned   := strings.Join(chunks, "\n")
    headed    := appendHeaders(aligned)

    err = ioutil.WriteFile(store.File, []byte(headed), 0644)
    if err != nil {
        return err
    }

    err = ioutil.WriteFile(store.File + ".json", data, 0644)
    if err != nil {
        return err
    }

    var loaded *Store
    err = json.Unmarshal(data, &loaded)
    if err != nil {
        return err
    }

    return nil
}

func (store *Store) SetItem(name string, value string) error {
    var item *Item
    item, err := NewItem(name, value)
    if err != nil {
        return err
    }

    store.Items[name] = *item

    return nil
}

func (store *Store) RenameItem(name string, new string) error {
    item, exists := store.Items[name]
    if !exists {
        // return errors.New(fmt.Sprintf("Item %s doesn't exist.", name))
        return fmt.Errorf("Item %s doesn't exist.", name)
    }

    // item.Name = new
    store.Items[new] = item
    delete(store.Items, name)

    return nil
}

func (store *Store) GetItem(name string) (*Item, error) {
    item, exists := store.Items[name]
    if !exists {
        // return nil, errors.New(fmt.Sprintf("Item %s doesn't exist.", name))
        return nil, fmt.Errorf("Item %s doesn't exist.", name)
    }

    return &item, nil
}

func (store *Store) DelItem(name string) error {
    _, exists := store.Items[name]
    if !exists {
        // return errors.New(fmt.Sprintf("Item %s doesn't exist.", name))
        return fmt.Errorf("Item %s doesn't exist.", name)
    }

    delete(store.Items, name);

    return nil
}

func (store *Store) AddTag(name string, tag string) error {
    item, err := store.GetItem(name)
    if err != nil {
        return fmt.Errorf("Item %s doesn't exist.", name)
    }

    if exists(item.Tags, tag) {
        return nil
    }

    item.Metadata.UpdatedAt = time.Now()
    item.Tags = append(item.Tags, tag)
    store.Items[name] = *item

    return nil
}

func (store *Store) DelTag(name string, tag string) bool {
    item, err := store.GetItem(name)
    if err != nil {
        return false
    }

    if !exists(item.Tags, tag) {
        return false
    }

    item.Metadata.UpdatedAt = time.Now()
    item.Tags = remove(item.Tags, tag)
    store.Items[name] = *item

    return true
}
