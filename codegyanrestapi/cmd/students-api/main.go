package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sahil3982/students-api/internal/config"
	"github.com/sahil3982/students-api/internal/http/handlers/student"
)

func main() {
	// Load configuration
	cfg := config.MustLoad()

	// Setup router and routes
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New() )

	// Create server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	// Start server
	slog.Info("starting the server...", slog.String("address", cfg.Addr))
	fmt.Printf("Server starting on %s...\n", cfg.Addr)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {

		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Failed to start server: %s", err)
		}
	}()
	<-done

	slog.Info("shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server stoopped successfully")
}
