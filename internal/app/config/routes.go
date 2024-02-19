package config

import (
	"kuchy/pages"
	"kuchy/pkg/dic"
	"net/http"
)

// RegisterRoutes registers routes
func RegisterRoutes(mux *http.ServeMux, c dic.Container) {
	page := c.Get("pages").(*pages.Pages)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if err := page.HomeIndex().Render(r.Context(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
