package handler

import (
	"github.com/codeluft/kuchy/templates/pages"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// Stock is the handler for the stock page.
func (h *handler) Stock(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	if err := pages.Stock().Render(h.ctx, w); err != nil {
		log.Fatal(err)
	}
}

// StockContents is the handler for the stock page contents.
func (h *handler) StockContents(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	if err := pages.StockContents().Render(h.ctx, w); err != nil {
		log.Fatal(err)
	}
}
