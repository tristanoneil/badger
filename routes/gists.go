package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/tristanoneil/badger/models"
)

func gists(w http.ResponseWriter, r *http.Request) {
	if !authorize(w, r) {
		return
	}

	user, _ := models.FindUser(currentUser(r).Username)
	gists := user.Gists()
	render("gists/index", w, r, map[string]interface{}{"Gists": gists})
}

func usersGists(w http.ResponseWriter, r *http.Request) {
	user, err := models.FindUser(mux.Vars(r)["username"])

	if err != nil {
		w.WriteHeader(404)
		return
	}

	gists := user.PublicGists()
	render("gists/index", w, r, map[string]interface{}{"Gists": gists})
}

func newGist(w http.ResponseWriter, r *http.Request) {
	if !authorize(w, r) {
		return
	}

	public, _ := strconv.ParseBool(r.FormValue("public"))

	gist := &models.Gist{
		UserID:    currentUser(r).ID,
		Title:     r.FormValue("title"),
		Content:   r.FormValue("content"),
		Public:    public,
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

func editGist(w http.ResponseWriter, r *http.Request) {
	if !authorize(w, r) {
		return
	}

	gist := models.FindGistForID(mux.Vars(r)["id"])

	if r.Method == "POST" {
		gist.Title = r.FormValue("title")
		gist.Content = r.FormValue("content")
		gist.UpdatedAt = time.Now()
		gist.Public, _ = strconv.ParseBool(r.FormValue("public"))

		if gist.Validate() {
			gist.Save()
			setSession(fmt.Sprintf("Successfully updated %s.", gist.Title), w, r)
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}

	render("gists/edit", w, r, map[string]interface{}{"Gist": gist})
}

func deleteGist(w http.ResponseWriter, r *http.Request) {
	if !authorize(w, r) {
		return
	}

	gist := models.FindGistForID(mux.Vars(r)["id"])
	gist.Delete()
	setSession(fmt.Sprintf("Successfully deleted %s.", gist.Title), w, r)
	http.Redirect(w, r, "/", http.StatusFound)
}
