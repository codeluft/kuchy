package handler

import (
	"context"
	"encoding/json"
	"github.com/a-h/templ"
	"github.com/codeluft/kuchy/view"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// Handler is a wrapper around the context, logger and translator function.
type Handler struct {
	ctx context.Context
	log *log.Logger
	tFn view.TranslatorFunc
}

// TemplateHandlerOpts is used to configure a TemplateHandler.
type TemplateHandlerOpts struct {
	pushUrl string
}

// TemplateHandlerOptsFunc is used to configure a TemplateHandler.
type TemplateHandlerOptsFunc func(t *TemplateHandlerOpts) *TemplateHandlerOpts

// New returns a new Handler.
func New(ctx context.Context, log *log.Logger, tFn view.TranslatorFunc) *Handler {
	return &Handler{ctx, log, tFn}
}

// JSONEncode encodes the given value as JSON and writes it to the response.
func (h *Handler) JSONEncode(w http.ResponseWriter, v interface{}) error {
	var jsonEncoder = json.NewEncoder(w)
	if w.Header().Get("Content-Type") != "application/json" {
		w.Header().Set("Content-Type", "application/json")
	}
	return jsonEncoder.Encode(v)
}

// JSONError writes the given error as JSON to the response.
func (h *Handler) JSONError(w http.ResponseWriter, err error, status int) {
	if w.Header().Get("Content-Type") != "application/json" {
		w.Header().Set("Content-Type", "application/json")
	}
	w.WriteHeader(status)
	h.log.Println(err)
	_ = h.JSONEncode(w, map[string]string{"error": err.Error()})
}

// TemplateHandler returns a httprouter.Handle that renders the given template.
func (h *Handler) TemplateHandler(template templ.Component, optsFuncs ...TemplateHandlerOptsFunc) httprouter.Handle {
	return func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		var opts = new(TemplateHandlerOpts)
		for _, fn := range optsFuncs {
			opts = fn(opts)
		}

		if opts.pushUrl != "" {
			w.Header().Add("HX-Push-Url", opts.pushUrl)
		}

		if err := template.Render(h.ctx, w); err != nil {
			log.Fatal(err)
		}
	}
}

// WithPushUrl returns a TemplateHandlerOptsFunc that sets the push url.
func WithPushUrl(url string) TemplateHandlerOptsFunc {
	return func(t *TemplateHandlerOpts) *TemplateHandlerOpts {
		t.pushUrl = url
		return t
	}
}

// TranslatorFunc returns the TranslatorFunc for the Handler.
func (h *Handler) TranslatorFunc() view.TranslatorFunc {
	return h.tFn
}
