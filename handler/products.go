package handler

import (
	"github.com/codeluft/kuchy/view/page"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// Products is the Handler for the products page.
func (h *Handler) Products(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	if err := page.Products(h.t).Render(h.ctx, w); err != nil {
		log.Fatal(err)
	}
}

// ProductsContents is the Handler for the products page contents.
func (h *Handler) ProductsContents(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	if err := page.ProductsContents(h.t).Render(h.ctx, w); err != nil {
		log.Fatal(err)
	}
}
