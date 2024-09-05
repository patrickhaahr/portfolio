package main

import (
	"GOTH/handlers"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()

	router.Handle("/*", public())
	router.Get("/", handlers.Make(handlers.HandleHome))
	router.Post("/toggle-icon", handlers.Make(handlers.ToggleIcon)) // Add this line for the toggle icon route
	router.Post("/submit-form", handlers.Make(handlers.HandleContactForm))
	router.Get("/close-modal", handlers.Make(handlers.CloseModal)) // Add this new line

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // default port if not set
	}
	log.Printf("Server starting on port %s", port)
	http.ListenAndServe(":"+port, router)
}
