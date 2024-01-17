package service

import (
	. "github.com/codeluft/kuchy/domain/model"
)

var (
	productServiceInstance *ProductService
)

type productRepository interface {
	CreateProduct(uuid, name, barcode string, measurementValue float64, measUnit *MeasurementUnitModel) (*ProductModel, error)
}

type measurementUnitRepository interface {
	GetMeasurementUnitByUuid(uuid string) (*MeasurementUnitModel, error)
}

type ProductService struct {
	productRepo  productRepository
	measUnitRepo measurementUnitRepository
}

// NewProduct creates a new ProductService instance.
func NewProduct(productRepo productRepository, measUnitRepo measurementUnitRepository) *ProductService {
	return &ProductService{productRepo: productRepo, measUnitRepo: measUnitRepo}
}

// SingletonProduct creates a new ProductService instance if it doesn't exist yet.
func SingletonProduct(productRepo productRepository, measUnitRepo measurementUnitRepository) *ProductService {
	if productServiceInstance == nil {
		productServiceInstance = NewProduct(productRepo, measUnitRepo)
	}
	return productServiceInstance
}

// CreateProduct creates a new product in the database.
func (s *ProductService) CreateProduct(uuid, name, barcode, measUnitUuid string, measValue float64) (*ProductModel, error) {
	measUnit, err := s.measUnitRepo.GetMeasurementUnitByUuid(measUnitUuid)
	if err != nil {
		return nil, err
	}

	return s.productRepo.CreateProduct(uuid, name, barcode, measValue, measUnit)
}