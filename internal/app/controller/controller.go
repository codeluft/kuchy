package controller

import (
	"net/http"
)

// Default is the default controller.
type Default struct {
}

// PushUrl sets the HX-Push-Url header.
func (d *Default) PushUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Push-Url", r.URL.Path)
}
