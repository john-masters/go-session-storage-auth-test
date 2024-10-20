package handlers

import (
	"fmt"
	"go-session-storage-auth-test/db"
	"go-session-storage-auth-test/models"
	"net/http"
	"time"

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
	fmt.Println("testing testing")

	userCount, err := db.SelectUserCountByEmail(email)
	if err != nil {
		fmt.Println("Error getting user count:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("user count", userCount)
	if userCount > 0 {
		fmt.Fprint(w, "An account with this email already exists")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error generating hash from password:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userId, err := db.InsertUser(&models.User{
		Email:    email,
		Password: string(hash),
	})

	// userId will be 0 if error
	if err != nil && userId != 0 {
		fmt.Println("Error creating user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sessionCount, err := db.SelectSessionCountByUserId(userId)

	if err != nil {
		fmt.Println("Error getting session count:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if sessionCount > 0 {
		err = db.DeleteSessionByUserId(userId)
		if err != nil {
			fmt.Println("Error deleting existing user sessions:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	sessionId, err := db.InsertSession(userId)
	if err != nil {
		fmt.Println("Error creating session:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "Session",
		Value:   sessionId,
		Expires: time.Now().Add(time.Hour * 24),
		Path:    "/",
	})
	w.Header().Add("HX-Redirect", "/account")
	w.WriteHeader(http.StatusCreated)
}
