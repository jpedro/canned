package can

import (
    "fmt"
    // "strings"
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

// func strip(text string) string {
//     return strings.Replace(text, "\n", "", -1)
// }

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

func appendHeaders(text string) string {
    headers := make(map[string]string)
    headers["version"] = CAN_VERSION
    headers["algorithm"] = CAN_ALGORITHM
    var header string
    for key, val := range headers {
        header = fmt.Sprintf("%s%s: %s\n", header, key, val)
    }

    return fmt.Sprintf("%s\n%s", header, text)
}

// func splitHeaders(text string) ([]string, string) {
//     headers := make(map[string]string)
//     headers["version"] = VERSION
//     var header string
//     for key, val := range headers {
//         header = fmt.Sprintf("%s%s: %s\n", header, key, val)
//     }

//     return fmt.Sprintf("%s\n%s", header, text), "a"
// }
