package controller

import (
	"context"
	"github.com/codeluft/kuchy/internal/app/translator"
	"log"
)

// Product is the controller for the product related pages.
type Product struct {
	Default
	ctx context.Context
	log *log.Logger
	tfn translator.Func
}

// NewProduct returns a new Product controller.
func NewProduct(ctx context.Context, log *log.Logger, tfn translator.Func) *Product {
	return &Product{
		ctx: ctx,
		log: log,
		tfn: tfn,
	}
}
