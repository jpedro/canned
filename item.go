package canned

type Item struct {
    // Name        string      `json:"name"`
    Content     string      `json:"content"`
    Metadata    Metadata    `json:"metadata"`
    Tags        []string    `json:"tags"`
}
