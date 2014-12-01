package main

import "net/http"

func signup(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"Email": r.FormValue("email")}

	if r.Method == "POST" {
		var count int
		db.Get(&count, `
			SELECT COUNT(*) FROM users WHERE email = $1
			`, r.FormValue("email"))

		if r.FormValue("password") != r.FormValue("password_confirmation") {
			setSession("Passwords don't match.", w, r)
		} else if count > 0 {
			setSession("Email must be unique.", w, r)
		} else {
			db.MustExec(`
				INSERT into users (email, password)
					VALUES ($1, crypt($2, gen_salt('bf')))
			`, r.FormValue("email"), r.FormValue("password"))

			setSession("Successfully signed up", w, r)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
	}

	render("signup", w, r, data)
}

func login(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"Email": r.FormValue("email")}

	if r.Method == "POST" {
		var count int
		db.Get(&count, `
			SELECT COUNT(*) FROM users WHERE email = $1
				AND password = crypt($2, password)
		`, r.FormValue("email"), r.FormValue("password"))

		if count > 0 {
			setSession(r.FormValue("email"), w, r, "Email")
			setSession("Successfully logged in.", w, r)

			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			setSession("Invalid username or password.", w, r)
		}
	}

	render("login", w, r, data)
}
