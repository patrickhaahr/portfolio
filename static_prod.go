//go:build !dev && !heroku
// +build !dev,!heroku

package main

import (
	"embed"
	"net/http"
)

//go:embed public
var publicFS embed.FS

func public() http.Handler {
	return http.FileServer(http.FS(publicFS))
}
