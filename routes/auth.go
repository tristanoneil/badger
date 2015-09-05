package routes

import (
	"net/http"

	"github.com/tristanoneil/badger/models"
)

func signup(w http.ResponseWriter, r *http.Request) {
	user := &models.User{
		Email:                r.FormValue("email"),
		Username:             r.FormValue("username"),
		Password:             r.FormValue("password"),
		PasswordConfirmation: r.FormValue("password_confirmation"),
	}

	if r.Method == "POST" {
		if usernameConflictsWithRoute(user.Username) {
			setSession("Username is a reserved word.", w, r, "Error")
		} else {
			if user.Validate() {
				user.Create()
				setSession("Successfully signed up.", w, r)
				setSession(user.Email, w, r, "Email")
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}
		}
	}

	renderTemplate("signup", w, r, map[string]interface{}{"User": user})
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

	renderTemplate("login", w, r, data)
}

func logout(w http.ResponseWriter, r *http.Request) {
	setSession(nil, w, r, "Email")
	setSession("Successfully logged out.", w, r)
	http.Redirect(w, r, "/login", http.StatusFound)
}
