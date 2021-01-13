package lib

import (
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

func strim(text string) string {
    return strings.Replace(text, "\n", "", -1)
}
