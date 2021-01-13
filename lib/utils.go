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
