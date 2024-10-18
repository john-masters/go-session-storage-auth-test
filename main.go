package main

import (
	"fmt"
	"go-session-storage-auth-test/db"
	"go-session-storage-auth-test/routes"
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
		http.ServeFile(w, r, "index.html")
	})

	mux.Handle("/auth/", http.StripPrefix("/auth", routes.AuthRoutes()))

	fmt.Printf("Server running on http://localhost:%v\n", port)

	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
