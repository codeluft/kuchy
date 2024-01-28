package main

import (
	"embed"
	"flag"
	"fmt"
	"github.com/codeluft/kuchy/internal/app"
	"github.com/codeluft/kuchy/internal/app/router"
	"log"
	"net/http"
	"time"
)

//go:embed static
var assets embed.FS

//go:embed translations/*.yaml
var translationsFS embed.FS

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

	var addr = fmt.Sprintf("%s:%d", host, port)
	var container = app.NewContainer(translationsFS, assets)

	if handler, err := router.NewHandler(container); err == nil {
		var server = http.Server{
			Addr:         addr,
			Handler:      handler,
			IdleTimeout:  5 * time.Second,
			ReadTimeout:  1 * time.Second,
			WriteTimeout: 2 * time.Second,
		}

		log.Printf("Running http server at http://%s", addr)
		log.Fatal(server.ListenAndServe())
	} else {
		log.Fatal(err)
	}
}
