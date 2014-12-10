package routes

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tristanoneil/badger/models"
)

func gists(w http.ResponseWriter, r *http.Request) {
	if !authorize(w, r) {
		return
	}

	gists := models.GetGistsForUserID(currentUser(r).ID)
	render("gists/index", w, r, map[string]interface{}{"Gists": gists})
}

func newGist(w http.ResponseWriter, r *http.Request) {
	if !authorize(w, r) {
		return
	}

	gist := &models.Gist{
		UserID:    currentUser(r).ID,
		Title:     r.FormValue("title"),
		Content:   r.FormValue("content"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if r.Method == "POST" && gist.Validate() {
		gist.Create()
		setSession("Successfully created gist.", w, r)
		http.Redirect(w, r, "/", http.StatusFound)
	}

	render("gists/new", w, r, map[string]interface{}{"Gist": gist})
}

func showGist(w http.ResponseWriter, r *http.Request) {
	if !authorize(w, r) {
		return
	}

	gist := models.FindGistForID(mux.Vars(r)["id"])
	render("gists/show", w, r, map[string]interface{}{"Gist": gist})
}
