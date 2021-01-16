package canned

// Item struct
type Item struct {
    Content     string      `json:"content" yaml:"content"`
    Metadata    Metadata    `json:"metadata" yaml:"metadata"`
    Tags        []string    `json:"tags,omitempty" yaml:"tags,omitempty"`
}
