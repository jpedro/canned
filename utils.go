package canned

import (
    "fmt"
    "strings"
)

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

func strip(text string) string {
    return strings.Replace(text, "\n", "", -1)
}

func exists(list []string, search string) bool {
    for _, val := range list {
        if val == search {
            return true
        }
    }

    return false
}

func remove(list []string, search string) ([]string) {
    for index, value := range list {
        if value == search {
            return append(list[:index], list[index + 1:]...)
        }
    }

    return list
}

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

func getHeaders(text string) (map[string]string, string) {
    index   := strings.Index(text, SEPARATOR)
    header  := text[0:index]
    headers := make(map[string]string)
    payload := text[index + len(SEPARATOR):]

    parts := strings.Split(header, "\n")
    for part := range parts {
        line := parts[part]
        // fmt.Printf("Part %d: %s.\n", part, line)
        colon := strings.Index(line, ":")
        key := line[0:colon]
        val := line[colon + 2:]
        // fmt.Printf("Key=%s Value=%s.\n", key, val)
        headers[key] = string(val)
    }

    return headers, payload
}

func verifyPassword(password string) error {
    if password == "" {
        return fmt.Errorf("Can password is required.")
    }

    return nil
}
