package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/codeluft/kuchy/pkg/app"
	"github.com/codeluft/kuchy/translations"
	"log"
	"net/http"
)

//go:embed static
var assets embed.FS

//go:embed translations/*.yaml
var transFS embed.FS

const (
	ServerPort = 3000
)

func main() {
	var ctx = context.TODO()
	var t, err = translations.NewTranslator(transFS)
	var session = app.NewSession()

	if err != nil {
		log.Fatal(err)
	}

	var router = app.NewRouter(assets, ctx, t, session)
	var addr = fmt.Sprintf(":%d", ServerPort)

	log.Printf("Running http server at %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
