package repository

import (
	. "github.com/codeluft/kuchy/domain/model"
	"github.com/lqs/sqlingo"
)

var (
	measurementUnitRepositoryInstance *MeasurementUnitRepository
)

type MeasurementUnitRepository struct {
	db sqlingo.Database
}

// NewMeasurementUnit creates a new MeasurementUnitRepository instance.
func NewMeasurementUnit(db sqlingo.Database) *MeasurementUnitRepository {
	return &MeasurementUnitRepository{db}
}

// SingletonMeasurementUnit creates a new MeasurementUnitRepository instance if it doesn't exist yet.
func SingletonMeasurementUnit(db sqlingo.Database) *MeasurementUnitRepository {
	if measurementUnitRepositoryInstance == nil {
		measurementUnitRepositoryInstance = NewMeasurementUnit(db)
	}
	return measurementUnitRepositoryInstance
}

// GetMeasurementUnitByUuid returns a measurement unit by its uuid.
func (r *MeasurementUnitRepository) GetMeasurementUnitByUuid(uuid string) (*MeasurementUnitModel, error) {
	var measUnit *MeasurementUnitModel
	_, err := r.db.SelectFrom(MeasurementUnit).
		Where(MeasurementUnit.Uuid.Equals(uuid)).
		FetchFirst(measUnit)
	if err != nil {
		return nil, err
	}

	return measUnit, nil
}
