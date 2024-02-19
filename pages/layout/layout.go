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
		return templ.Raw(fmt.Sprintf(
			`<script type="text/javascript">console.error("Error loading script %s: %w")</script>`,
			path,
			err,
		))
	}

	return templ.Raw(fmt.Sprintf(`<script type="text/javascript">%s</script>`, string(contents)))
}
