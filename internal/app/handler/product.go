package handler

import (
	"github.com/codeluft/kuchy/internal/domain/model"
	"github.com/codeluft/kuchy/internal/view/page"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// ProductsListing renders the products listing page.
func (h *Handler) ProductsListing(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	var products []model.Product
	if err := page.Products(h.tFn, products).Render(h.ctx, w); err != nil {
		log.Fatal(err)
	}
}

// ProductsListingContents renders the products listing page.
func (h *Handler) ProductsListingContents(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	var products []model.Product
	w.Header().Add("HX-Push-Url", "/products")
	if err := page.ProductsContents(h.tFn, products).Render(h.ctx, w); err != nil {
		log.Fatal(err)
	}
}

// ProductsAdd renders the products listing page.
func (h *Handler) ProductsAdd(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	if err := page.ProductsAdd(h.tFn).Render(h.ctx, w); err != nil {
		log.Fatal(err)
	}
}

// ProductsAddContents renders the products listing page.
func (h *Handler) ProductsAddContents(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	if err := page.ProductsAdd(h.tFn).Render(h.ctx, w); err != nil {
		log.Fatal(err)
	}
}
