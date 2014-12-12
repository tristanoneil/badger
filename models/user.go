package models

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
)

//
// User is used to map users in the database.
//
type User struct {
	ID                   int    `db:"id"`
	Email                string `db:"email"`
	Username             string `db:"username"`
	Password             string `db:"password"`
	PasswordConfirmation string
	Errors               map[string]string
}

//
// Validate determines if a given User is valid.
//
func (user *User) Validate() bool {
	user.Errors = make(map[string]string)

	if !user.UniqueEmail() {
		user.Errors["Email"] = "Email must be unique."
	}

	if user.Email == "" {
		user.Errors["Email"] = "You must provide an email."
	}

	if user.Username == "" {
		user.Errors["Username"] = "You must provide a username."
	}

	if !user.UniqueUsername() {
		user.Errors["Username"] = "Username must be unique."
	}

	if user.Password == "" {
		user.Errors["Password"] = "You must provide a password."
	}

	if user.Password != user.PasswordConfirmation {
		user.Errors["Password"] = "Passwords must match."
	}

	return len(user.Errors) == 0
}

//
// GravatarURL returns a Gravatar URL for a given user and size.
//
func (user User) GravatarURL(size int) string {
	h := md5.New()
	io.WriteString(h, user.Email)
	return fmt.Sprintf("https://secure.gravatar.com/avatar/%x?s=%d", h.Sum(nil), size)
}

//
// Create creates a new user in the database.
//
func (user *User) Create() {
	_, err := Db.NamedExec(
		`INSERT into users (email, username, password)
		VALUES (:email, :username, crypt(:password, gen_salt('bf')))`, user,
	)

	if err != nil {
		log.Fatal(err)
	}
}

//
// UniqueEmail determines if a given user is unique based on their email.
//
func (user *User) UniqueEmail() bool {
	var count int
	err := Db.Get(&count, "SELECT COUNT(*) FROM users WHERE email = $1", user.Email)

	if err != nil {
		log.Fatal(err)
	}

	return count == 0
}

//
// UniqueUsername determines if a given user is unique based on their email.
//
func (user *User) UniqueUsername() bool {
	var count int
	err := Db.Get(&count, "SELECT COUNT(*) FROM users WHERE username = $1", user.Username)

	if err != nil {
		log.Fatal(err)
	}

	return count == 0
}

//
// Gists returns all gists for a user oredered by the created at date.
//
func (user *User) Gists() []Gist {
	gists := []Gist{}
	err := Db.Select(
		&gists,
		`SELECT * FROM gists WHERE user_id = $1 ORDER BY created_at DESC`,
		user.ID,
	)

	if err != nil {
		log.Fatal(err)
	}

	return gists
}

//
// IsValidUser determines if a given email address and password
// are a valid user.
//
func IsValidUser(email string, password string) bool {
	var count int
	err := Db.Get(
		&count, `SELECT COUNT(*) FROM users WHERE email = $1
		AND password = crypt($2, password)`, email, password,
	)

	if err != nil {
		log.Fatal(err)
	}

	return count > 0
}

//
// FindUser returns a user for a given identifier.
//
func FindUser(identifier string) (User, error) {
	var user User
	err := Db.Get(&user, "SELECT * FROM users WHERE email = $1 OR username = $1", identifier)
	return user, err
}
