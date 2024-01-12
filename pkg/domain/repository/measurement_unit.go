package repository

import (
	. "github.com/codeluft/kuchy/pkg/domain/model"
	"github.com/lqs/sqlingo"
)

var (
	measurementUnitRepositoryInstance *measurementUnitRepository
)

type measurementUnitRepository struct {
	db sqlingo.Database
}

// NewMeasurementUnit creates a new measurementUnitRepository instance.
func NewMeasurementUnit(db sqlingo.Database) *measurementUnitRepository {
	return &measurementUnitRepository{db}
}

// SingletonMeasurementUnit creates a new measurementUnitRepository instance if it doesn't exist yet.
func SingletonMeasurementUnit(db sqlingo.Database) *measurementUnitRepository {
	if measurementUnitRepositoryInstance == nil {
		measurementUnitRepositoryInstance = NewMeasurementUnit(db)
	}
	return measurementUnitRepositoryInstance
}

// GetMeasurementUnitByUuid returns a measurement unit by its uuid.
func (r *measurementUnitRepository) GetMeasurementUnitByUuid(uuid string) (*MeasurementUnitModel, error) {
	var measUnit *MeasurementUnitModel
	_, err := r.db.SelectFrom(MeasurementUnit).
		Where(MeasurementUnit.Uuid.Equals(uuid)).
		FetchFirst(measUnit)
	if err != nil {
		return nil, err
	}

	return measUnit, nil
}