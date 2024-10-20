package routes

import (
	"go-session-storage-auth-test/handlers"
	"go-session-storage-auth-test/middleware"
	"net/http"
)

func PageRoutes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /", handlers.Home)
	router.HandleFunc("GET /account", middleware.RequireAuth(handlers.Account))

	return router
}
