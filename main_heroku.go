//go:build heroku
// +build heroku

package main

import (
	"log"
	"net/http"
	"os"

	"GOTH/handlers"

	"github.com/go-chi/chi/v5"
)

func public() http.Handler {
	return http.FileServer(http.Dir("public"))
}

func main() {
	router := chi.NewMux()

	router.Handle("/*", public())
	router.Get("/", handlers.Make(handlers.HandleHome))
	router.Post("/toggle-icon", handlers.Make(handlers.ToggleIcon))
	router.Post("/submit-form", handlers.Make(handlers.HandleContactForm))
	router.Get("/close-modal", handlers.Make(handlers.CloseModal))

	port := os.Getenv("PORT")
	log.Printf("Server starting on port %s", port)
	http.ListenAndServe(":"+port, router)
}
