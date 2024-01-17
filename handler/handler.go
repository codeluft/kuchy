package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	ctx context.Context
	log *log.Logger
	t   func(string, string) string
}

// New returns a new Handler.
func New(ctx context.Context, l *log.Logger, t func(string, string) string) *Handler {
	return &Handler{ctx: ctx, log: l, t: t}
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
