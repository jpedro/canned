package canned

import (
	"time"
)

// Metadata struct
type Metadata struct {
	CreatedAt time.Time `json:"createdAt" yaml:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" yaml:"updatedAt"`
}
