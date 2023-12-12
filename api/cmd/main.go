package main

import (
	"log"
	"net/http"

	"example.com/internal/api"
	"example.com/internal/config"
	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
)

func main() {
    // Load configuration (environment variables, etc.)

    cfg := config.Load()
    

    // Initialize the API handlers
    router := api.NewRouter()

    // Start the server
    log.Println("Starting server on port", cfg.Port)
    log.Fatal(http.ListenAndServe(":"+cfg.Port, handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),
        handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
    )(router)))
}
