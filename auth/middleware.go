package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
)

const (
	username     = "admin"
	passwordHash = "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918" // SHA256 hash of "admin"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := r.Cookie("session")
		if session == nil || session.Value != username {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Login(username, password string) bool {
	hash := sha256.Sum256([]byte(password))
	return username == username && hex.EncodeToString(hash[:]) == passwordHash
}
