package table

import (
	"time"
)

// Base model
type Base struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
