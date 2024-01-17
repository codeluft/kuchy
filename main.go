package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"github.com/codeluft/kuchy/app"
	"log"
	"net/http"
	"time"
)

//go:embed static
var assets embed.FS

//go:embed translations/*.yaml
var transFS embed.FS

const (
	ServerHost = "localhost"
	ServerPort = 42069
)

func main() {
	var host string
	var port int

	flag.StringVar(&host, "host", ServerHost, "The host to listen on.")
	flag.IntVar(&port, "port", ServerPort, "The port to listen on.")
	flag.Parse()

	translator, err := app.NewTranslator(transFS)
	if err != nil {
		log.Fatal(err)
	}

	var serverHandler = app.NewServerHandler().
		WithLogger(log.Default()).
		WithContext(context.TODO()).
		WithAssets(assets).
		WithSessionManager(app.NewSessionManager()).
		WithTranslator(translator).
		Register()

	var addr = fmt.Sprintf("%s:%d", host, port)
	var server = http.Server{
		Addr:         addr,
		Handler:      serverHandler,
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	log.Printf("Running http server at %s", addr)
	log.Fatal(server.ListenAndServe())
}
