package app

import (
	"github.com/codeluft/kuchy/app/handler"
	"github.com/codeluft/kuchy/view/page"
	"github.com/julienschmidt/httprouter"
)

func routes(r *httprouter.Router, h *handler.Handler) {
	var tFn = h.TranslatorFunc()

	// Home routes
	r.GET("/", h.TemplateHandler(page.Home(tFn)))
	r.GET("/pages/home", h.TemplateHandler(page.HomeContents(tFn)))

	// Stock routes
	r.GET("/stock", h.TemplateHandler(page.Stock(tFn)))
	r.GET("/pages/stock", h.TemplateHandler(page.StockContents(tFn)))

	// Recipes routes
	r.GET("/recipes", h.TemplateHandler(page.Recipes(tFn)))
	r.GET("/pages/recipes", h.TemplateHandler(page.RecipesContents(tFn)))

	// Products routes
	r.GET("/products", h.ProductsListing)
	r.GET("/pages/products", h.ProductsListingContents)
	r.GET("/products/add", h.ProductsAdd)
	r.GET("/pages/products/add", h.ProductsAddContents)
}
