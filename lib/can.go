package lib

import (
    "fmt"
    "time"
    // "errors"
    "strings"
    "encoding/json"
    "io/ioutil"
)

type Metadata struct {
    CreatedAt   time.Time   `json:"createdAt"`
    UpdatedAt   time.Time   `json:"updatedAt"`
}

type Can struct {
    File        string              `json:"-"`
    Version     string              `json:"version"`
    Metadata    Metadata            `json:"metadata"`
    Items       map[string]Item     `json:"items"`
}

func Open(file string) (*Can, error) {
    // fmt.Printf("Opening %v...\n", can.File)
    can := &Can{}
    can.File = file
    err := can.Parse()

    return can, err
}

func (can *Can) Parse() error {
    data, err := ioutil.ReadFile(can.File + ".json")
    if err != nil {
        // return errors.New("Couldn't read json file.")
        return err
    }

    err = json.Unmarshal(data, &can)
    if err != nil {
        // return errors.New("Couldn't decoded can file.")
        return err
    }

    return nil
}

func NewCan(file string) (*Can, error) {
    can := &Can{}
    can.File = file
    can.Version = "v1"
    can.Metadata.CreatedAt = time.Now()
    can.Items = make(map[string]Item)

    return can, nil
}

func Init(file string) (*Can, error) {
    can, err := NewCan(file)
    if err != nil {
        return nil, err
    }

    err = can.Save()

    return can, nil
}

func Create(file string) (*Can, error) {
    can, err := NewCan(file)

    return can, err
}

func (can *Can) Save() error {
    // fmt.Printf("Saving %v...\n", can.File)
    data, err := json.MarshalIndent(can, "", "  ")
    if err != nil {
        // return errors.New("Couldn't encode can.")
        return err
    }

    encrypted := Encrypt(string(data), "pp")
    chunks    := align(encrypted, 64)
    aligned   := strings.Join(chunks, "\n")
    strimed   := strim(aligned)
    headed    := addHeaders(strimed)

    // fmt.Println(aligned)
    // fmt.Println(strimed)
    // fmt.Println(headed)

    err = ioutil.WriteFile(can.File, []byte(headed), 0644)
    if err != nil {
        // return errors.New("Couldn't write can file.")
        return err
    }

    err = ioutil.WriteFile(can.File + ".json", data, 0644)
    if err != nil {
        // return errors.New("Couldn't write json file.")
        return err
    }

    var loaded *Can
    err = json.Unmarshal(data, &loaded)
    if err != nil {
        // return errors.New("Couldn't decoded can file.")
        return err
    }

    // fmt.Printf("Original : %v\n", can)
    // fmt.Printf("Decoded  : %v\n", loaded)
    // fmt.Printf("Done\n")
    return nil
}

func addHeaders(text string) string {
    headers := make(map[string]string)
    headers["version"] = "v1"
    var header string
    for key, val := range headers {
        header = fmt.Sprintf("%s%s: %s\n", header, key, val)
    }

    return fmt.Sprintf("%s\n%s", header, text)
}

func chopHeaders(text string) (string, string) {
    headers := make(map[string]string)
    headers["version"] = "v1"
    var header string
    for key, val := range headers {
        header = fmt.Sprintf("%s%s: %s\n", header, key, val)
    }

    return fmt.Sprintf("%s\n%s", header, text), "a"
}

func (can *Can) SetItem(name string, value string) error {
    var item *Item
    item, err := NewItem(name, value)
    can.Items[name] = *item
    // fmt.Println(can.Items)

    return err
}
