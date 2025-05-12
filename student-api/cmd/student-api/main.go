package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"os/signal"

	"github.com/blue-samarth/trying-go/student-api/internal/config"
)

func main(){
	fmt.Println("Hello, World!")
	fmt.Println("This is a simple Go program.")
	cfg := config.MustLoad()
	fmt.Printf("Loaded config:\n  Env: %s\n  Storage Path: %s\n  HTTP Address: %s\n",
		cfg.Env, cfg.Storagepath, cfg.HTTPServer.Address)

	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Student API!"))
	})

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

}