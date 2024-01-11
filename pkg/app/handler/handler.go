package handler

import "context"

type handler struct {
	ctx context.Context
}

// New returns a new handler.
func New(ctx context.Context) *handler {
	return &handler{ctx: ctx}
}
