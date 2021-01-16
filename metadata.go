package can

import (
    "time"
)

type Metadata struct {
    CreatedAt   time.Time   `json:"createdAt"`
    UpdatedAt   time.Time   `json:"updatedAt"`
}
