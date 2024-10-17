package main

import (
	"fmt"
	"go-session-storage-auth-test/db"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Init()

	mux := http.NewServeMux()

	port := os.Getenv("PORT")

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "yo yo yo")
	})

	fmt.Printf("Server running on http://localhost:%v\n", port)

	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
