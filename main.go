package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/codeluft/kuchy/app"
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
	t, err := app.NewTranslator(transFS)
	if err != nil {
		log.Fatal(err)
	}

	var router = app.NewRouter(assets, context.TODO(), t, app.NewSessionManager())
	var addr = fmt.Sprintf(":%d", ServerPort)

	log.Printf("Running http server at %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
