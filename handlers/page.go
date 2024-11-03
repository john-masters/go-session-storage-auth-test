package handlers

import "net/http"

func HomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func SignupPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/signup.html")
}

func AccountPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/account.html")
}
