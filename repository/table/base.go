package table

import (
	"time"

	"github.com/uptrace/bun"
)

// Base model
type Base struct {
	bun.BaseModel
	ID        string     `bun:",pk"`
	CreatedAt *time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt bun.NullTime
	DeletedAt bun.NullTime `bun:",soft_delete,nullzero"`
}
