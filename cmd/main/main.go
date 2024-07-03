package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func serverRun(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "main")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading%v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not .env file")
	}


	http.HandleFunc("/", serverRun)
	fmt.Printf("server is listening on port %s...\n", port)
	fmt.Printf("http://localhost:%s\n", port)


	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v ", err)
	}

	
}
