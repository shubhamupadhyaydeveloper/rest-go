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

	"github.com/shubhamupadhyaydeveloper/students-api/internal/config"
	"github.com/shubhamupadhyaydeveloper/students-api/internal/http/handlers/student"
)

func main() {
	// load config
	cfg := config.Load()

	// setup Router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to go server created by shubham Upadhyay"))
	})
	router.HandleFunc("POST /users", student.CreateUserHandler)

	// setup Server
	server := http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: router,
	}

	slog.Info("Server is running on", slog.Int("port", cfg.HTTPServer.PORT))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt,syscall.SIGINT,syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}

	}()

	<-done

	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("error while shutting down server",slog.String("error",err.Error()))
	}

	slog.Info("server shutdown successful")
}
