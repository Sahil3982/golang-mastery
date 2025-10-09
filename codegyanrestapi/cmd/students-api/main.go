package main

import (
	"fmt"
	"github.com/sahil3982/students-api/internal/config"
	"log"
	"net/http"
)

func main() {
	// Load configuration
	cfg := config.MustLoad()

	// Setup router and routes
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Create server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router, 
	}

	// Start server
	fmt.Printf("Server starting on %s...\n", cfg.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
