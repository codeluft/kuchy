package app

import (
	"context"
	"embed"
	"github.com/codeluft/kuchy/pkg/app/handler"
	"github.com/codeluft/kuchy/templates"
	"github.com/codeluft/kuchy/translations"
	"github.com/julienschmidt/httprouter"
	"io/fs"
	"log"
	"net/http"
)

func registerHandlers(r *httprouter.Router, ctx context.Context, l *log.Logger, t translations.TranslatorFunc) {
	var h = handler.New(ctx, l, t)

	r.GET("/", h.Home)
	r.GET("/pages/home", h.HomeContents)

	r.GET("/stock", h.Stock)
	r.GET("/pages/stock", h.StockContents)

	r.GET("/recipes", h.Recipes)
	r.GET("/pages/recipes", h.RecipesContents)
}

type translator interface {
	Translate(string, string) string
	SetLanguage(string) error
}

type loggerRouter struct {
	log    *log.Logger
	router *httprouter.Router
	t      translator
}

// NewRouter returns a new http.Handler that serves the application.
func NewRouter(assets embed.FS, ctx context.Context, t translator) *loggerRouter {
	var httpRouter = httprouter.New()
	var logRouter = &loggerRouter{
		log:    log.Default(),
		router: httpRouter,
		t:      t,
	}

	staticFS, err := fs.Sub(assets, "static")
	if err != nil {
		log.Fatal(err)
	}
	httpRouter.ServeFiles("/static/*filepath", http.FS(staticFS))
	httpRouter.NotFound = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		if err := templates.NotFound().Render(ctx, w); err != nil {
			log.Fatal(err)
		}
	})

	registerHandlers(httpRouter, ctx, logRouter.log, logRouter.t.Translate)

	return logRouter
}

// ServeHTTP implements http.Handler with logging.
func (r *loggerRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	lang := req.URL.Query().Get("lang")
	if lang == "" {
		lang = translations.DefaultLanguage
	}

	err := r.t.SetLanguage(lang)
	if err != nil {
		r.log.Println(err)
	}

	r.log.Printf("%s %s", req.Method, req.URL.Path)
	r.router.ServeHTTP(w, req)
}
