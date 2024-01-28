package controller

import (
	"context"
	"github.com/codeluft/kuchy/internal/app/translator"
	"github.com/codeluft/kuchy/internal/view/page"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// Home is the controller for the home page.
type Home struct {
	Default
	ctx context.Context
	log *log.Logger
	tfn translator.Func
}

// NewHome returns a new Home controller.
func NewHome(ctx context.Context, log *log.Logger, tfn translator.Func) *Home {
	return &Home{
		ctx: ctx,
		log: log,
		tfn: tfn,
	}
}

// Index is the handler for the home page.
func (h *Home) Index(resp http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var component = page.Home
	if contents := req.URL.Query().Get("contents"); contents == "true" {
		component = page.HomeContents
	}
	h.PushUrl(resp, req)

	if err := component(h.tfn).Render(h.ctx, resp); err != nil {
		h.log.Fatal(err)
	}
}
