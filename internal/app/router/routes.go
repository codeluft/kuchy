package router

import (
	"github.com/codeluft/kuchy/internal/app"
	"github.com/julienschmidt/httprouter"
)

func Routes(r *httprouter.Router, c *app.Container) {
	r.GET("/", c.Controller.Home.Index)
}
