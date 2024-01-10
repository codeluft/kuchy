package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/codeluft/kuchy-frontend/templates"
	"github.com/codeluft/kuchy-frontend/templates/pages"
	"github.com/julienschmidt/httprouter"
	"io/fs"
	"log"
	"net/http"
)

//go:embed static
var assets embed.FS

func main() {
	var router = httprouter.New()
	var ctx = context.TODO()
	var port = 3000
	var addr = fmt.Sprintf(":%d", port)

	staticFS, err := fs.Sub(assets, "static")
	if err != nil {
		log.Fatal(err)
	}
	router.ServeFiles("/static/*filepath", http.FS(staticFS))

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if err := templates.Home().Render(ctx, w); err != nil {
			log.Fatal(err)
		}
	})

	router.GET("/pages/home", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if err := pages.Home().Render(ctx, w); err != nil {
			log.Fatal(err)
		}
	})

	router.GET("/stock", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if err := templates.Stock().Render(ctx, w); err != nil {
			log.Fatal(err)
		}
	})

	router.GET("/pages/stock", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if err := pages.Stock().Render(ctx, w); err != nil {
			log.Fatal(err)
		}
	})

	log.Printf("Running http server at %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
