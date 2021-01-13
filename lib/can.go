package lib

import (
    "fmt"
    "time"
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

func Open(file string) *Can {
    can := &Can{}
    can.File = file
    fmt.Printf("Opening %v...\n", can.File)

    return can
}

func Init(file string) *Can {
    can := Create(file)
    can.Metadata.CreatedAt = time.Now()
    item := NewItem("test", "123")
    can.Items = make(map[string]Item)
    can.Items["test"] = *item
    can.Save()

    return can
}

func Create(file string) *Can {
    can := &Can{}
    can.File = file
    can.Version = "v1"
    fmt.Printf("Creating %v...\n", can.File)
    return can
}

func (can *Can) Save() {
    fmt.Printf("Saving %v...\n", can.File)

	data, err := json.MarshalIndent(can, "", "  ")
    if err != nil {
        panic("Couldn't encode can")
	}

	encrypted := Encrypt(string(data), "pp")
	chunks    := align(encrypted, 64)
	aligned   := strings.Join(chunks, "\n")
	unaligned := unalign(aligned)
	fmt.Println(aligned)
	fmt.Println(unaligned)

	err = ioutil.WriteFile(can.File, []byte(encrypted), 0644)
    if err != nil {
        panic("Couldn't write can file.")
    }

    err = ioutil.WriteFile(can.File + ".json", data, 0644)
    if err != nil {
        panic("Couldn't write json file.")
    }

    var loaded *Can
    err = json.Unmarshal(data, &loaded)
    if err != nil {
        panic("Couldn't decoded can file.")
    }

	fmt.Printf("Original : %v\n", can)
    fmt.Printf("Decoded  : %v\n", loaded)
    fmt.Printf("Done\n")
}

func align(text string, size int) []string {
	var chunks []string
	runes := []rune(text)

	if len(runes) == 0 {
		return []string{text}
	}

	for i := 0; i < len(runes); i += size {
		nn := i + size
		if nn > len(runes) {
			nn = len(runes)
		}
		chunks = append(chunks, string(runes[i:nn]))
	}

	return chunks
}

func unalign(text string) string {
	return strings.Replace(text, "\n", "", -1)
}
