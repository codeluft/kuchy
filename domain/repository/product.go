package repository

import (
	. "github.com/codeluft/kuchy/domain/model"
	"github.com/lqs/sqlingo"
	"time"
)

type ProductRepository struct {
	db sqlingo.Database
}

// NewProduct creates a new ProductRepository instance.
func NewProduct(db sqlingo.Database) *ProductRepository {
	return &ProductRepository{db}
}

// CreateProduct creates a new product in the database.
func (r *ProductRepository) CreateProduct(uuid, name, barcode string, measurementValue float64, measUnit *MeasurementUnitModel) (*ProductModel, error) {
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
