package models

import (
	"html/template"
	"log"
	"time"

	"github.com/russross/blackfriday"
)

//
// Gist is used to map gists in the database.
//
type Gist struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Errors    map[string]string
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
	_, err := db.NamedExec(`
		INSERT into gists (title, content, user_id, created_at, updated_at)
		VALUES (:title, :content, :user_id, :created_at, :updated_at)`, gist,
	)

	if err != nil {
		log.Fatal(err)
	}
}

//
// Markdown returns rendered markdown from the gists Content.
//
func (gist Gist) Markdown() template.HTML {
	return template.HTML(string(blackfriday.MarkdownCommon([]byte(gist.Content))))
}

//
// GetGistsForUserID Gets all gists for a given user id.
//
func GetGistsForUserID(UserID int) []Gist {
	gists := []Gist{}
	err := db.Select(
		&gists,
		`SELECT * FROM gists WHERE user_id = $1 ORDER BY created_at DESC`,
		UserID,
	)

	if err != nil {
		log.Fatal(err)
	}

	return gists
}

func FindGist(ID interface{}) Gist {
	gist := Gist{}
	db.Get(&gist, "SELECT * FROM gists WHERE id = $1", ID)
	return gist
}
