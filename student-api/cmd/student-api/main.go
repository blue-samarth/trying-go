package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"os/signal"
	"log/slog"
	"syscall"
	"context"
	"time"

	"github.com/blue-samarth/trying-go/student-api/internal/config"
	"github.com/blue-samarth/trying-go/student-api/cmd/http/handlers/student"
)

func main(){
	fmt.Println("Hello, World!")
	fmt.Println("This is a simple Go program.")
	cfg := config.MustLoad()
	fmt.Printf("Loaded config:\n  Env: %s\n  Storage Path: %s\n  HTTP Address: %s\n",
		cfg.Env, cfg.Storagepath, cfg.HTTPServer.Address)

	router := http.NewServeMux()
	router.HandleFunc("GET /api/students", student.New())

	fmt.Println("Server started on", cfg.HTTPServer.Address)
	fmt.Println("Server is running...")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	server := http.Server {
		Addr : cfg.HTTPServer.Address,
		Handler : router,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil { log.Fatalf("Error starting server: %v", err.Error()) }
	}()

	<-done 
	fmt.Println("Shutting down server...")
	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	err := server.Shutdown(ctx) // Gracefully shutdown the server but issue here is that sometimes while shutting down the server, it will not shutdown gracefully and will throw an error
	if err != nil { slog.Error("Error shutting down server", slog.String("error", err.Error())) }
	fmt.Println("Server stopped")

	slog.Info("Server stopped successfully")

}