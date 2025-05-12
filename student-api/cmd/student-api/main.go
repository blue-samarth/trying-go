package main

import (
	"fmt"
	"github.com/blue-samarth/trying-go/student-api/internal/config"
)

func main(){
	fmt.Println("Hello, World!")
	fmt.Println("This is a simple Go program.")
	cfg := config.MustLoad()
	fmt.Printf("Loaded config:\n  Env: %s\n  Storage Path: %s\n  HTTP Address: %s\n",
		cfg.Env, cfg.Storagepath, cfg.HTTPServer.Address)
}