package models

//
// User is used to map users in the database.
//
type User struct {
	ID                   int    `db:"id"`
	Email                string `db:"email"`
	Password             string `db:"password"`
	PasswordConfirmation string
	Errors               map[string]string
}

//
// Validate determines if a given User is valid.
//
func (user *User) Validate() bool {
	user.Errors = make(map[string]string)

	if !IsUniqueUser(user.Email) {
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

//
// Create creates a new user in the database.
//
func (user *User) Create() {
	db.NamedExec(`INSERT into users (email, password) VALUES (:email, crypt(:password, gen_salt('bf')))`, user)
}

//
// IsUniqueUser determines if a given email address is unique.
//
func IsUniqueUser(email string) bool {
	var count int
	db.Get(&count, "SELECT COUNT(*) FROM users WHERE email = $1", email)
	return count == 0
}

//
// IsValidUser determines if a given email address and password are a valid user.
//
func IsValidUser(email string, password string) bool {
	var count int
	db.Get(&count, `SELECT COUNT(*) FROM users WHERE email = $1 AND password = crypt($2, password)`, email, password)
	return count > 0
}

//
// GetUserIDForEmail returns a users id for a given email address.
//
func GetUserIDForEmail(email string) int {
	var ID int
	db.Get(&ID, "SELECT id FROM users WHERE email = $1", email)
	return ID
}
