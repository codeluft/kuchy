package handler

import (
	"github.com/codeluft/kuchy/view/page"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// Recipes is the Handler for the recipes page.
func (h *Handler) Recipes(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	if err := page.Recipes(h.t).Render(h.ctx, w); err != nil {
		log.Fatal(err)
	}
}

// RecipesContents is the Handler for the recipes page contents.
func (h *Handler) RecipesContents(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	if err := page.RecipesContents(h.t).Render(h.ctx, w); err != nil {
		log.Fatal(err)
	}
}
