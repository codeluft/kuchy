package layout

import (
	"embed"
	"fmt"
	"github.com/a-h/templ"
)

type Layout struct {
	translator translator
	staticFs   *embed.FS
}

type translator interface {
	Translate(string) string
}

// New creates a new Layout
func New(t translator, fs *embed.FS) *Layout {
	return &Layout{
		translator: t,
		staticFs:   fs,
	}
}

// Translate translates a string
func (l *Layout) Translate(v string) string {
	return l.translator.Translate(v)
}

// InlineScript returns a script tag with the contents of the file at the given path.
func (l *Layout) InlineScript(path string) templ.Component {
	contents, err := l.staticFs.ReadFile(path)
	if err != nil {
		return l.script([]byte(fmt.Sprintf(`console.error("Error loading script %s: %s")`, path, err)))
	}

	return l.script(contents)
}

// InlineStyle returns a style tag with the contents of the file at the given path.
func (l *Layout) InlineStyle(path string) templ.Component {
	contents, err := l.staticFs.ReadFile(path)
	if err != nil {
		return l.style([]byte(fmt.Sprintf(`/* Error loading style %s: %s */`, path, err)))
	}

	return l.style(contents)
}

func (l *Layout) script(contents []byte) templ.Component {
	return templ.Raw(fmt.Sprintf(`<script type="text/javascript" async>%s</script>`, string(contents)))
}

func (l *Layout) style(contents []byte) templ.Component {
	return templ.Raw(fmt.Sprintf(`<style type="text/css">%s</style>`, string(contents)))
}
