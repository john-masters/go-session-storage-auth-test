package routes

import (
	"go-session-storage-auth-test/handlers"
	"net/http"
)

func PageRoutes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /", handlers.Home)
	router.HandleFunc("GET /account", handlers.Account)

	return router
}
