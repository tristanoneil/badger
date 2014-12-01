package main

import "net/http"

type User struct {
	ID                   int    `db:"id"`
	Email                string `db:"email"`
	Password             string `db:"password"`
	PasswordConfirmation string
	Errors               map[string]string
}

func (user *User) Validate() bool {
	user.Errors = make(map[string]string)

	var count int
	db.Get(&count, `
		SELECT COUNT(*) FROM users WHERE email = $1
		`, user.Email)

	if count > 0 {
		user.Errors["Email"] = "Email must be unique."
	}

	if user.Email == "" {
		user.Errors["Email"] = "You must provide an email."
	}

	if user.Password == "" {
		user.Errors["Password"] = "You must provide a password."
	}

	if user.Password != user.PasswordConfirmation {
		user.Errors["Password"] = "Passwords must match."
	}

	return len(user.Errors) == 0
}

func signup(w http.ResponseWriter, r *http.Request) {
	user := &User{
		Email:                r.FormValue("email"),
		Password:             r.FormValue("password"),
		PasswordConfirmation: r.FormValue("password_confirmation"),
	}

	if r.Method == "POST" && user.Validate() {
		db.NamedExec(`
				INSERT into users (email, password)
					VALUES (:email, crypt(:password, gen_salt('bf')))
			`, user)

		setSession("Successfully signed up.", w, r)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	render("signup", w, r, map[string]interface{}{"User": user})
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
