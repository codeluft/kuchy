package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/codeluft/kuchy/pkg/app"
	"log"
	"net/http"
)

//go:embed static
var assets embed.FS

const (
	ServerPort = 3000
)

func main() {
	var ctx = context.TODO()
	var router = app.NewRouter(assets, ctx)
	var addr = fmt.Sprintf(":%d", ServerPort)

	log.Printf("Running http server at %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
