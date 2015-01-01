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
	Public    bool      `db:"public"`
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
	_, err := Db.NamedExec(`
		INSERT into gists (title, content, public, user_id, created_at, updated_at)
		VALUES (:title, :content, :public, :user_id, :created_at, :updated_at)`, gist,
	)

	if err != nil {
		log.Fatal(err)
	}
}

//
// Delete deletes a gist in the database.
//
func (gist *Gist) Delete() {
	_, err := Db.NamedExec("DELETE FROM gists WHERE id = :id", gist)

	if err != nil {
		log.Fatal(err)
	}
}

//
// Save saves a gists properties in the database.
//
func (gist *Gist) Save() {
	_, err := Db.NamedExec(`
		UPDATE gists SET title = :title,
		content = :content, public = :public, updated_at = :updated_at WHERE id = :id`, gist,
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
// FindGistForID returns a Gist for a given ID.
//
func FindGistForID(ID interface{}) Gist {
	gist := Gist{}
	Db.Get(&gist, "SELECT * FROM gists WHERE id = $1", ID)
	return gist
}
