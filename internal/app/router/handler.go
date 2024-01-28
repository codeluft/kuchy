package router

import (
	"errors"
	"github.com/codeluft/kuchy/internal/app"
	"github.com/codeluft/kuchy/internal/app/translator"
	"github.com/codeluft/kuchy/internal/view/layout"
	"github.com/julienschmidt/httprouter"
	"io/fs"
	"log"
	"net/http"
)

var (
	ErrNilContainer = errors.New("container is nil")
)

// Handler is a http.Handler that serves the application.
type Handler struct {
	router    *httprouter.Router
	container *app.Container
}

// NewHandler returns a new Handler that serves the application.
func NewHandler(c *app.Container) (*Handler, error) {
	if c != nil {
		var handler = &Handler{
			router:    httprouter.New(),
			container: c,
		}

		// Setup static file handler.
		staticFS, err := fs.Sub(c.Assets, "static")
		if err != nil {
			log.Fatal(err)
		}
		handler.router.ServeFiles("/static/*filepath", http.FS(staticFS))

		// Setup 404 handler.
		handler.router.NotFound = http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			resp.Header().Set("Content-Type", "text/html; charset=utf-8")
			resp.WriteHeader(http.StatusNotFound)
			if err := layout.NotFound().Render(c.Context, resp); err != nil {
				log.Fatal(err)
			}
		})

		// Setup language switch handler.
		handler.router.GET("/lang/:lang", func(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
			c.Session.Manager.GetSession(resp, req).Set("lang", params.ByName("lang"))
			http.Redirect(resp, req, "/", http.StatusFound)
		})

		// Load routes.
		Routes(handler.router, c)

		return handler, nil
	}

	return nil, ErrNilContainer
}

// ServeHTTP implements http.Handler with logging.
func (h *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if c := h.container; c != nil {
		lang := c.Session.Manager.GetSession(resp, req).Get("lang")
		if lang == nil {
			lang = translator.DefaultLanguage
		}
		c.Session.Manager.GetSession(resp, req).Set("lang", lang)

		if err := c.Translator.Loader.SetLanguage(lang.(string)); err != nil {
			c.Log.Println(err)
		}

		c.Log.Printf("%s %s", req.Method, req.URL.Path)
		h.router.ServeHTTP(resp, req)
	}
}
