package app

import (
	"context"
	"github.com/codeluft/kuchy/handler"
	"github.com/codeluft/kuchy/view/layout"
	"github.com/julienschmidt/httprouter"
	"io/fs"
	"log"
	"net/http"
)

type translatorFunc func(string, string) string

func registerHandlers(r *httprouter.Router, ctx context.Context, l *log.Logger, t translatorFunc) {
	var h = handler.New(ctx, l, t)

	r.GET("/", h.Home)
	r.GET("/pages/home", h.HomeContents)

	r.GET("/stock", h.Stock)
	r.GET("/pages/stock", h.StockContents)

	r.GET("/recipes", h.Recipes)
	r.GET("/pages/recipes", h.RecipesContents)

	r.GET("/products", h.Products)
	r.GET("/pages/products", h.ProductsContents)
}

type Translator interface {
	Translate(string, string) string
	SetLanguage(string) error
}

type FileSystem interface {
	fs.ReadDirFS
	fs.ReadFileFS
}

type router struct {
	router *httprouter.Router
	log    *log.Logger
	t      Translator
	s      *SessionManager
}

// NewRouter returns a new http.Handler that serves the application.
func NewRouter(assets FileSystem, ctx context.Context, t Translator, s *SessionManager) *router {
	var appRouter = &router{httprouter.New(), log.Default(), t, s}

	staticFS, err := fs.Sub(assets, "static")
	if err != nil {
		log.Fatal(err)
	}
	appRouter.router.ServeFiles("/static/*filepath", http.FS(staticFS))
	appRouter.router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		if err := layout.NotFound().Render(ctx, w); err != nil {
			log.Fatal(err)
		}
	})

	appRouter.router.GET("/lang/:lang", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		appRouter.s.GetSession(w, r).Set("lang", ps.ByName("lang"))
		http.Redirect(w, r, "/", http.StatusFound)
	})

	registerHandlers(appRouter.router, ctx, appRouter.log, appRouter.t.Translate)

	return appRouter
}

// ServeHTTP implements http.Handler with logging.
func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	lang := r.s.GetSession(w, req).Get("lang")
	if lang == nil {
		lang = DefaultLanguage
	}
	r.s.GetSession(w, req).Set("lang", lang)

	if err := r.t.SetLanguage(lang.(string)); err != nil {
		r.log.Println(err)
	}

	r.log.Printf("%s %s", req.Method, req.URL.Path)
	r.router.ServeHTTP(w, req)
}
