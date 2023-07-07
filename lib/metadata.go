package canned

import (
	"time"
)

// Metadata struct
type Metadata struct {
	Created time.Time `json:"created" yaml:"created"`
	Updated time.Time `json:"updated,omitempty" yaml:"updated,omitempty"`
}
