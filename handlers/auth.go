package handlers

import (
	"fmt"
	"go-session-storage-auth-test/db"
	"go-session-storage-auth-test/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Println("Error parsing form")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	switch {
	case email == "":
		fmt.Fprint(w, "Email is required")
		return
	case password == "":
		fmt.Fprint(w, "Password is required")
		return
	}

	var userCount int

	err = db.SelectUserCountByEmail(email, &userCount)
	if err != nil {
		fmt.Println("Error getting user count:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if userCount > 0 {
		fmt.Fprint(w, "An account with this email already exists")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error generating hash from password:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = db.InsertUser(&models.User{
		Email:    email,
		Password: string(hash),
	})
	if err != nil {
		fmt.Println("Error creating user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}
