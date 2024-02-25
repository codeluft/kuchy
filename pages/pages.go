package pages

import (
	"embed"
	"kuchy/pages/layout"
)

type translatorAware interface {
	Translate(string) string
}

// Pages is a collection of pages
type Pages struct {
	translator translatorAware
	staticFs   *embed.FS
	layout     *layout.Layout
}

// New creates a new Pages
func New(t translatorAware, fs *embed.FS, l *layout.Layout) *Pages {
	return &Pages{
		translator: t,
		staticFs:   fs,
		layout:     l,
	}
}

// Translate translates a string
func (p *Pages) Translate(v string) string {
	return p.translator.Translate(v)
}
