package app

import (
	"context"
	"embed"
	"github.com/codeluft/kuchy/pkg/app/handler"
	"github.com/codeluft/kuchy/templates"
	"github.com/julienschmidt/httprouter"
	"io/fs"
	"log"
	"net/http"
)

func registerHandlers(r *httprouter.Router, ctx context.Context, l *log.Logger) {
	var h = handler.New(ctx, l)

	r.GET("/", h.Home)
	r.GET("/pages/home", h.HomeContents)

	r.GET("/stock", h.Stock)
	r.GET("/pages/stock", h.StockContents)

	r.GET("/recipes", h.Recipes)
	r.GET("/pages/recipes", h.RecipesContents)
}

type loggerRouter struct {
	log    *log.Logger
	router *httprouter.Router
}

// NewRouter returns a new http.Handler that serves the application.
func NewRouter(assets embed.FS, ctx context.Context) http.Handler {
	var router = httprouter.New()
	var defaultLog = log.Default()

	staticFS, err := fs.Sub(assets, "static")
	if err != nil {
		log.Fatal(err)
	}
	router.ServeFiles("/static/*filepath", http.FS(staticFS))
	registerHandlers(router, ctx, defaultLog)

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		if err := templates.NotFound().Render(ctx, w); err != nil {
			log.Fatal(err)
		}
	})

	return &loggerRouter{
		log:    defaultLog,
		router: router,
	}
}

// ServeHTTP implements http.Handler with logging.
func (r *loggerRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.log.Printf("%s %s", req.Method, req.URL.Path)
	r.router.ServeHTTP(w, req)
}
