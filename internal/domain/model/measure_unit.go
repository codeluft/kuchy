package model

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MeasurementUnit struct {
	bun.BaseModel `bun:"table:measurement_unit"`

	ID                 int64 `bun:",pk,autoincrement"`
	UUID               uuid.UUID
	ReferenceUnit      *MeasurementUnit
	ReferenceUnitValue float64
	Name               string
	CreatedAt          bun.NullTime `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt          bun.NullTime `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt          bun.NullTime `bun:",nullzero"`
}
