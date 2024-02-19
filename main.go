package main

import (
	"embed"
	"kuchy/internal/app"
)

//go:embed static
var staticFs embed.FS

func main() {
	app.New(&staticFs).Run(":8080")
}
