package canned

import (
    "fmt"
    "strings"
)

// Splits a text into an aligned column of text
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

// Strips a string from line linux endings
func strip(text string) string {
    return strings.Replace(text, "\n", "", -1)
}

// Checks if an element exists in a list
func exists(list []string, search string) bool {
    for _, val := range list {
        if val == search {
            return true
        }
    }

    return false
}

// Removes an element from a list
func remove(list []string, element string) ([]string) {
    for index, value := range list {
        if value == element {
            return append(list[:index], list[index + 1:]...)
        }
    }

    return list
}

// Appends headers to a text
func addHeaders(text string) string {
    headers := make(map[string]string)
    headers["version"] = VERSION
    headers["algorithm"] = ALGORITHM
    var header string
    for key, val := range headers {
        header = fmt.Sprintf("%s%s: %s\n", header, key, val)
    }

    return fmt.Sprintf("%s\n%s", header, text)
}

// Parses the headers and returns them and the payload
func getHeaders(text string) (map[string]string, string) {
    index   := strings.Index(text, SEPARATOR)
    header  := text[0:index]
    headers := make(map[string]string)
    payload := text[index + len(SEPARATOR):]

    parts := strings.Split(header, "\n")
    for part  := range parts {
        line  := parts[part]
        colon := strings.Index(line, ":")
        key   := line[0:colon]
        value := line[colon + 2:]
        headers[key] = string(value)
    }

    return headers, payload
}

// Ensures the password is not empty or errors
func verifyPassword(password string) error {
    if password == "" {
        return fmt.Errorf("Password cannot be empty.")
    }

    return nil
}
