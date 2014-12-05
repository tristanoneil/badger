package main

import (
	"net/http"

	"github.com/tristanoneil/badger/models"
)

func signup(w http.ResponseWriter, r *http.Request) {
	user := &models.User{
		Email:                r.FormValue("email"),
		Password:             r.FormValue("password"),
		PasswordConfirmation: r.FormValue("password_confirmation"),
	}

	if r.Method == "POST" && user.Validate() {
		user.Create()
		setSession("Successfully signed up.", w, r)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	render("signup", w, r, map[string]interface{}{"User": user})
}

func login(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"Email": r.FormValue("email")}

	if r.Method == "POST" {
		if models.IsValidUser(r.FormValue("email"), r.FormValue("password")) {
			setSession(r.FormValue("email"), w, r, "Email")
			setSession("Successfully logged in.", w, r)

			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			setSession("Invalid username or password.", w, r)
		}
	}

	render("login", w, r, data)
}
