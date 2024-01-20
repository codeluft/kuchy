package app

import (
	"github.com/codeluft/kuchy/view/page"
	"github.com/julienschmidt/httprouter"
)

func routes(r *httprouter.Router, h *Handler) {
	var tFn = h.TranslatorFunc()

	r.GET("/", h.TemplateHandler(page.Home(tFn)))
	r.GET("/pages/home", h.TemplateHandler(page.HomeContents(tFn)))

	r.GET("/stock", h.TemplateHandler(page.Stock(tFn)))
	r.GET("/pages/stock", h.TemplateHandler(page.StockContents(tFn)))

	r.GET("/recipes", h.TemplateHandler(page.Recipes(tFn)))
	r.GET("/pages/recipes", h.TemplateHandler(page.RecipesContents(tFn)))

	r.GET("/products", h.TemplateHandler(page.Products(tFn)))
	r.GET("/pages/products", h.TemplateHandler(page.ProductsContents(tFn)))
	r.GET("/pages/products/add", h.TemplateHandler(page.ProductsAddContents(tFn)))
}
