package middleware

import (
	"fmt"
	"go-session-storage-auth-test/db"
	"net/http"
)

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("Session")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		sessionCount, err := db.SelectSessionCountById(session.Value)

		if err != nil {
			fmt.Println("Error getting session count:", err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		if sessionCount > 0 {
			next(w, r)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
}
