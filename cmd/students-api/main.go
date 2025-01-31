package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mehulsuthar-000/students-api/internal/config"
	"github.com/mehulsuthar-000/students-api/internal/http/handlers/student"
	"github.com/mehulsuthar-000/students-api/internal/storage/sqlite"
)

func main() {
	// load the configuration
	cfg := config.MustLoad()

	// connect to the database
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	slog.Info("Storage initialized", slog.String("env", cfg.ENV), slog.String("storage", "sqlite"))

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))

	// setup server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	log.Printf("server started on %s", cfg.Address)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("server failed to start: %v", err)
		}
	}()

	<-done
	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("server failed to shutdown: ", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")
}
