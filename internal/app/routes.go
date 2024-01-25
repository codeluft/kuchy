package app

import (
	. "github.com/codeluft/kuchy/internal/app/handler"
	"github.com/codeluft/kuchy/internal/view/page"
	"github.com/julienschmidt/httprouter"
)

func routes(r *httprouter.Router, h *Handler) {
	var tFn = h.TranslatorFunc()

	// Home routes
	r.GET("/", h.TemplateHandler(page.Home(tFn)))
	r.GET("/pages/home", h.TemplateHandler(page.HomeContents(tFn), WithPushUrl("/")))

	// Stock routes
	r.GET("/stock", h.TemplateHandler(page.Stock(tFn)))
	r.GET("/pages/stock", h.TemplateHandler(page.StockContents(tFn), WithPushUrl("/stock")))

	// Recipes routes
	r.GET("/recipes", h.TemplateHandler(page.Recipes(tFn)))
	r.GET("/pages/recipes", h.TemplateHandler(page.RecipesContents(tFn), WithPushUrl("/recipes")))

	// Products routes
	r.GET("/products", h.ProductsListing)
	r.GET("/pages/products", h.ProductsListingContents)
	r.GET("/products/add", h.ProductsAdd)
	r.GET("/pages/products/add", h.ProductsAddContents)
}
