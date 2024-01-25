package model

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Product struct {
	bun.BaseModel `bun:"table:product"`

	ID                   int64 `bun:",pk,autoincrement"`
	UUID                 uuid.UUID
	MeasurementUnit      *MeasurementUnit
	MeasurementUnitValue float64
	Name                 string
	Barcode              string
	CreatedAt            bun.NullTime `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt            bun.NullTime `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt            bun.NullTime `bun:",nullzero"`
}
