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
			data["Flash"] = "Passwords don't match."
		} else if count > 0 {
			data["Flash"] = "Email must be unique."
		} else {
			db.MustExec(`
				INSERT into users (email, password)
					VALUES ($1, crypt($2, gen_salt('bf')))
			`, r.FormValue("email"), r.FormValue("password"))
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
			session, _ := store.Get(r, "auth")
			session.Values["Email"] = r.FormValue("email")
			session.Values["Flash"] = "Successfully logged in."
			session.Save(r, w)

			http.Redirect(w, r, "/gists", http.StatusFound)
		} else {
			data["Flash"] = "Invalid username or password."
		}
	}

	render("login", w, r, data)
}
