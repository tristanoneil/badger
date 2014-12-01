package main

import "net/http"

type Gist struct {
	ID      int    `db:"id"`
	UserID  int    `db:"user_id"`
	Title   string `db:"title"`
	Content string `db:"content"`
	Errors  map[string]string
}

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

func gists(w http.ResponseWriter, r *http.Request) {
	authorize(w, r)

	gists := []Gist{}
	db.Select(&gists, "SELECT * FROM gists WHERE user_id = $1", currentUserID(r))

	render("gists/index", w, r, map[string]interface{}{"Gists": gists})
}

func newGist(w http.ResponseWriter, r *http.Request) {
	authorize(w, r)

	gist := &Gist{
		UserID:  currentUserID(r),
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	}

	if r.Method == "POST" && gist.Validate() {
		db.NamedExec(`
				INSERT into gists (title, content, user_id)
				VALUES (:title, :content, :user_id)
			`, gist)

		setSession("Successfully created gist.", w, r)
		http.Redirect(w, r, "/", http.StatusFound)
	}

	render("gists/new", w, r, map[string]interface{}{"Gist": gist})
}
