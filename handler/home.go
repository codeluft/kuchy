package handler

import (
	"github.com/codeluft/kuchy/view/page"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// Home is the Handler for the home page.
func (h *Handler) Home(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	if err := page.Home(h.t).Render(h.ctx, w); err != nil {
		log.Fatal(err)
	}
}

// HomeContents is the Handler for the home page contents.
func (h *Handler) HomeContents(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	if err := page.HomeContents(h.t).Render(h.ctx, w); err != nil {
		log.Fatal(err)
	}
}
