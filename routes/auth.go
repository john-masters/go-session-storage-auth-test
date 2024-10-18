package routes

import (
	"go-session-storage-auth-test/handlers"
	"net/http"
)

func AuthRoutes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /sign-up", handlers.Signup)

	return router
}
