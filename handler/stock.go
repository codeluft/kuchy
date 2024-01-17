package handler

import (
	"github.com/codeluft/kuchy/view/page"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// Stock is the Handler for the stock page.
func (h *Handler) Stock(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	if err := page.Stock(h.t).Render(h.ctx, w); err != nil {
		log.Fatal(err)
	}
}

// StockContents is the Handler for the stock page contents.
func (h *Handler) StockContents(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	if err := page.StockContents(h.t).Render(h.ctx, w); err != nil {
		log.Fatal(err)
	}
}
