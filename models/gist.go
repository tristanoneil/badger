package models

import "log"

//
// Gist is used to map gists in the database.
//
type Gist struct {
	ID      int    `db:"id"`
	UserID  int    `db:"user_id"`
	Title   string `db:"title"`
	Content string `db:"content"`
	Errors  map[string]string
}

//
// Validate determines if a given Gist is valid.
//
func (gist *Gist) Validate() bool {
	gist.Errors = make(map[string]string)

	if gist.Title == "" {
		gist.Errors["Title"] = "You must provide a title."
	}

	if gist.Content == "" {
		gist.Errors["Content"] = "You must provide content."
	}

	return len(gist.Errors) == 0
}

//
// Create creates a new gist in the database.
//
func (gist *Gist) Create() {
	_, err := db.NamedExec(`INSERT into gists (title, content, user_id) VALUES (:title, :content, :user_id)`, gist)

	if err != nil {
		log.Fatal(err)
	}
}

//
// GetGistsForUserID Gets all gists for a given user id.
//
func GetGistsForUserID(UserID int) []Gist {
	gists := []Gist{}
	err := db.Select(&gists, "SELECT * FROM gists WHERE user_id = $1", UserID)

	if err != nil {
		log.Fatal(err)
	}

	return gists
}
