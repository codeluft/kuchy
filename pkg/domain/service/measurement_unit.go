package service

import (
	. "github.com/codeluft/kuchy/pkg/domain/model"
)

type measurementUnitService struct {
	measurementUnitRepo measurementUnitRepository
}

// NewMeasurementUnit creates a new measurementUnitService instance.
func NewMeasurementUnit(measurementUnitRepo measurementUnitRepository) *measurementUnitService {
	return &measurementUnitService{measurementUnitRepo}
}

// GetMeasurementUnitByUuid returns a measurement unit by its uuid.
func (s *measurementUnitService) GetMeasurementUnitByUuid(uuid string) (*MeasurementUnitModel, error) {
	return s.measurementUnitRepo.GetMeasurementUnitByUuid(uuid)
}
