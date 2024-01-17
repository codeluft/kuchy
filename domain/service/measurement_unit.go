package service

import (
	. "github.com/codeluft/kuchy/domain/model"
)

type MeasurementUnitService struct {
	measurementUnitRepo measurementUnitRepository
}

// NewMeasurementUnit creates a new MeasurementUnitService instance.
func NewMeasurementUnit(measurementUnitRepo measurementUnitRepository) *MeasurementUnitService {
	return &MeasurementUnitService{measurementUnitRepo}
}

// GetMeasurementUnitByUuid returns a measurement unit by its uuid.
func (s *MeasurementUnitService) GetMeasurementUnitByUuid(uuid string) (*MeasurementUnitModel, error) {
	return s.measurementUnitRepo.GetMeasurementUnitByUuid(uuid)
}
