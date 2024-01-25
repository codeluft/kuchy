package controller

import (
	c "context"
	t "github.com/codeluft/kuchy/internal/app/translator"
	"log"
)

var (
	context    *c.Context
	logger     *log.Logger
	translator *t.Loader
)

type Controller struct {
}

func (c *Controller) Init(ctx *c.Context, log *log.Logger, t *t.Loader) {
	if context == nil {
		context = ctx
	}

	if logger == nil {
		logger = log
	}

	if translator == nil {
		translator = t
	}
}

func (c *Controller) Context() *c.Context {
	return context
}
