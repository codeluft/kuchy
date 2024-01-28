package controller

import (
	"net/http"
)

// Default is a wrapper around the context, logger and translator function.
type Default struct {
}

// PushUrl sets the HX-Push-Url header.
func (d *Default) PushUrl(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("HX-Push-Url", req.URL.Path)
}
