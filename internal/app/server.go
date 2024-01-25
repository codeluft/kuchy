package app

import (
	"context"
	"errors"
	"github.com/codeluft/kuchy/internal/app/handler"
	"github.com/codeluft/kuchy/internal/app/session"
	"github.com/codeluft/kuchy/internal/app/translator"
	"github.com/codeluft/kuchy/internal/view/layout"
	"github.com/julienschmidt/httprouter"
	"io/fs"
	"log"
	"net/http"
)

var (
	ErrNilLogger         = errors.New("logger is nil")
	ErrNilSessionManager = errors.New("session manager is nil")
	ErrNilTranslator     = errors.New("translator is nil")
)

// FileSystem defines a contract for serving static files.
type FileSystem interface {
	fs.ReadDirFS
	fs.ReadFileFS
}

// ServerHandler is a http.Handler that serves the application.
type ServerHandler struct {
	router         *httprouter.Router
	log            *log.Logger
	translator     *translator.Loader
	sessionManager *session.Manager
	ctx            context.Context
}

// NewServerHandler returns a new ServerHandler that serves the application.
func NewServerHandler() *ServerHandler {
	return &ServerHandler{router: httprouter.New()}
}

// Register registers the handlers.
func (sh *ServerHandler) Register() *ServerHandler {
	sh.requireLogger()
	sh.requireContext()

	sh.router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		if err := layout.NotFound().Render(sh.ctx, w); err != nil {
			log.Fatal(err)
		}
	})
	routes(sh.router, handler.New(sh.ctx, sh.log, sh.translator.Translate))
	return sh
}

func (sh *ServerHandler) WithLogger(l *log.Logger) *ServerHandler {
	sh.log = l
	return sh
}

func (sh *ServerHandler) WithContext(ctx context.Context) *ServerHandler {
	sh.ctx = ctx
	return sh
}

func (sh *ServerHandler) WithSessionManager(sm *session.Manager) *ServerHandler {
	sh.sessionManager = sm
	return sh
}

func (sh *ServerHandler) WithTranslator(t *translator.Loader) *ServerHandler {
	sh.requireSessionManager()
	sh.translator = t
	sh.router.GET("/lang/:lang", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		sh.sessionManager.GetSession(w, r).Set("lang", ps.ByName("lang"))
		http.Redirect(w, r, "/", http.StatusFound)
	})
	return sh
}

func (sh *ServerHandler) WithAssets(assets FileSystem) *ServerHandler {
	sh.requireLogger()
	staticFS, err := fs.Sub(assets, "static")
	if err != nil {
		log.Fatal(err)
	}
	sh.router.ServeFiles("/static/*filepath", http.FS(staticFS))
	return sh
}

// ServeHTTP implements http.Handler with logging.
func (sh *ServerHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	sh.requireSessionManager()
	sh.requireTranslator()

	lang := sh.sessionManager.GetSession(w, req).Get("lang")
	if lang == nil {
		lang = translator.DefaultLanguage
	}
	sh.sessionManager.GetSession(w, req).Set("lang", lang)

	if err := sh.translator.SetLanguage(lang.(string)); err != nil {
		sh.log.Println(err)
	}

	sh.log.Printf("%s %s", req.Method, req.URL.Path)
	sh.router.ServeHTTP(w, req)
}

func (sh *ServerHandler) requireLogger() {
	if sh.log == nil {
		log.Default().Fatalln(ErrNilLogger)
	}
}

func (sh *ServerHandler) requireContext() {
	sh.requireLogger()
	if sh.ctx == nil {
		sh.log.Fatalln(ErrNilLogger)
	}
}

func (sh *ServerHandler) requireSessionManager() {
	sh.requireLogger()
	if sh.sessionManager == nil {
		sh.log.Fatalln(ErrNilSessionManager)
	}
}

func (sh *ServerHandler) requireTranslator() {
	sh.requireLogger()
	if sh.translator == nil {
		sh.log.Fatalln(ErrNilTranslator)
	}
}