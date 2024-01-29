package controller

import (
	"context"
	"github.com/codeluft/kuchy/internal/app/translator"
	"github.com/codeluft/kuchy/internal/domain/model"
	"github.com/codeluft/kuchy/internal/view/page"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
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

// Index is the handler for the product listing page.
func (p *Product) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var products []model.Product
	var component = page.Products
	if contents := r.URL.Query().Get("contents"); contents == "true" {
		component = page.ProductsContents
	}
	p.PushUrl(w, r)

	if err := component(p.tfn, products).Render(p.ctx, w); err != nil {
		p.log.Fatal(err)
	}
}
