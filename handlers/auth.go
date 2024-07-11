package handlers

import (
	"net/http"

	"PagProp/auth"
	"PagProp/views/login"
)

func HandleLoginForm(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, login.LoginForm())
}

func HandleLogin(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if auth.Login(username, password) {
		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: username,
			Path:  "/",
		})
		http.Redirect(w, r, "/uploadPropiedad", http.StatusSeeOther)
		return nil
	}

	return Render(w, r, login.LoginForm())
}

func HandleLogout(w http.ResponseWriter, r *http.Request) error {
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}
