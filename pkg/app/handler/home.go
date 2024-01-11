package handler

import (
	"github.com/codeluft/kuchy/templates/pages"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// Home is the handler for the home page.
func (h *handler) Home(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	if err := pages.Home().Render(h.ctx, w); err != nil {
		log.Fatal(err)
	}
}

// HomeContents is the handler for the home page contents.
func (h *handler) HomeContents(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	if err := pages.HomeContents().Render(h.ctx, w); err != nil {
		log.Fatal(err)
	}
}
