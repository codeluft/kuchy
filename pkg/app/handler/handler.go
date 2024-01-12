package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type handler struct {
	ctx context.Context
	log *log.Logger
}

// New returns a new handler.
func New(ctx context.Context, l *log.Logger) *handler {
	return &handler{ctx: ctx, log: l}
}

// JSONEncode encodes the given value as JSON and writes it to the response.
func (h *handler) JSONEncode(w http.ResponseWriter, v interface{}) error {
	var jsonEncoder = json.NewEncoder(w)
	if w.Header().Get("Content-Type") != "application/json" {
		w.Header().Set("Content-Type", "application/json")
	}
	return jsonEncoder.Encode(v)
}

// JSONError writes the given error as JSON to the response.
func (h *handler) JSONError(w http.ResponseWriter, err error, status int) {
	if w.Header().Get("Content-Type") != "application/json" {
		w.Header().Set("Content-Type", "application/json")
	}
	w.WriteHeader(status)
	h.log.Println(err)
	_ = h.JSONEncode(w, map[string]string{"error": err.Error()})
}
