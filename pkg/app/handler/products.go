package handler

import (
	"github.com/codeluft/kuchy/templates/pages"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// Products is the handler for the products page.
func (h *handler) Products(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	if err := pages.Products(h.t).Render(h.ctx, w); err != nil {
		log.Fatal(err)
	}
}

// ProductsContents is the handler for the products page contents.
func (h *handler) ProductsContents(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	if err := pages.ProductsContents(h.t).Render(h.ctx, w); err != nil {
		log.Fatal(err)
	}
}
