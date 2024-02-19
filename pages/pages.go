package pages

import (
	"embed"
	"fmt"
	"github.com/a-h/templ"
	"kuchy/pages/layout"
)

type translatorAware interface {
	Translate(string) string
}

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

// InlineScript returns a script tag with the contents of the file at the given path.
func (p *Pages) InlineScript(path string) templ.Component {
	contents, err := p.staticFs.ReadFile(path)
	if err != nil {
		return templ.Raw(fmt.Sprintf(
			`<script type="text/javascript">console.error("Error loading script %s: %w")</script>`,
			path,
			err,
		))
	}

	return templ.Raw(fmt.Sprintf(`<script type="text/javascript">%s</script>`, string(contents)))
}
