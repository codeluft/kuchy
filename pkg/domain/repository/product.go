package repository

import (
	. "github.com/codeluft/kuchy/pkg/domain/model"
	"github.com/lqs/sqlingo"
	"time"
)

type productRepository struct {
	db sqlingo.Database
}

// NewProduct creates a new productRepository instance.
func NewProduct(db sqlingo.Database) *productRepository {
	return &productRepository{db}
}

// CreateProduct creates a new product in the database.
func (r *productRepository) CreateProduct(uuid, name, barcode string, measurementValue float64, measUnit *MeasurementUnitModel) (*ProductModel, error) {
	_, err := r.db.InsertInto(Product).
		Fields(
			Product.Uuid,
			Product.Name,
			Product.Barcode,
			Product.MeasurementValue,
			Product.MeasurementUnitId,
			Product.CreatedAt,
			Product.UpdatedAt,
		).
		Values(uuid, name, barcode, measurementValue, measUnit.Id, time.Now(), time.Now()).
		Execute()
	if err != nil {
		return nil, err
	}

	var product *ProductModel
	_, err = r.db.SelectFrom(Product).
		Where(Product.Uuid.Equals(uuid)).
		FetchFirst(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
