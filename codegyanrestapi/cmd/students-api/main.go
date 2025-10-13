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
	"github.com/sahil3982/students-api/internal/storage/sqlite"
)

func main() {
	// Load configuration
	cfg := config.MustLoad()

	//Database connection

	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage Intialized", slog.String("env", cfg.Env), slog.String("storage_path", cfg.StoragePath), slog.String("version", "1.0.0"))

	// Setup router and routes
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))

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
