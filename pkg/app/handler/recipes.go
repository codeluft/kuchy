package handler

import (
	"github.com/codeluft/kuchy/templates/pages"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// Recipes is the handler for the recipes page.
func (h *handler) Recipes(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	if err := pages.Recipes(h.t).Render(h.ctx, w); err != nil {
		log.Fatal(err)
	}
}

// RecipesContents is the handler for the recipes page contents.
func (h *handler) RecipesContents(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	if err := pages.RecipesContents(h.t).Render(h.ctx, w); err != nil {
		log.Fatal(err)
	}
}
