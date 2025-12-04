package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shubhamupadhyaydeveloper/students-api/internal/config"
)

func main(){
	// load config
    cfg := config.Load()
    
	// setup Router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to go server created by shubham Upadhyay"))
	})

	// setup Server
	server := http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: router,
	}

	fmt.Println("Server is running on port", cfg.HTTPServer.PORT)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start the server")
	}

}

