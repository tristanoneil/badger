package main

import "net/http"

type Gist struct {
	ID      int    `db:"id"`
	UserID  int    `db:"user_id"`
	Title   string `db:"title"`
	Content string `db:"content"`
}

func gists(w http.ResponseWriter, r *http.Request) {
	authorize(w, r)
	gists := []Gist{}
	db.Select(&gists, "SELECT * FROM gists WHERE user_id = $1", currentUserID(r))
	render("gists.html", w, r, map[string]interface{}{"Gists": gists})
}
