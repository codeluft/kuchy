package app

import (
	"context"
	"embed"
	"github.com/codeluft/kuchy/internal/app/controller"
	"github.com/codeluft/kuchy/internal/app/session"
	"github.com/codeluft/kuchy/internal/app/translator"
	"log"
)

// Container is the dependency injection container.
type Container struct {
	Context    context.Context
	Log        *log.Logger
	Assets     embed.FS
	Translator struct {
		Loader *translator.Loader
	}
	Session struct {
		Manager *session.Manager
	}
	Controller struct {
		Home    *controller.Home
		Product *controller.Product
	}
}

// InitAware is a contract for initializing a controller.
type InitAware interface {
	Init(ctx context.Context, log *log.Logger, tFn translator.Func)
}

// NewContainer returns a new Container.
func NewContainer(tfs, assets embed.FS) *Container {
	var c = new(Container)

	loader, err := translator.NewLoader(tfs)
	if err != nil {
		log.Fatal(err)
	}

	c.Context = context.Background()
	c.Log = log.Default()
	c.Translator.Loader = loader
	c.Session.Manager = session.NewManager()
	c.Assets = assets

	// Controllers
	c.Controller.Home = controller.NewHome(c.Context, c.Log, c.Translator.Loader.Translate)
	c.Controller.Product = controller.NewProduct(c.Context, c.Log, c.Translator.Loader.Translate)

	return c
}
