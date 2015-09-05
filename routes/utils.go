package routes

import (
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/justinas/nosurf"
	"github.com/tristanoneil/badger/models"
	"github.com/unrolled/render"
)

var (
	store *sessions.CookieStore
)

func init() {
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
}

func renderTemplate(templateName string, w http.ResponseWriter,
	r *http.Request, binding map[string]interface{}) {

	session, _ := store.Get(r, "auth")

	if flash, ok := session.Values["Flash"].(string); ok {
		binding["Flash"] = flash
		setSession("", w, r)
	}

	if flash, ok := session.Values["Error"].(string); ok {
		binding["Error"] = flash
		setSession("", w, r, "Error")
	}

	binding["Token"] = nosurf.Token(r)

	if loggedIn(r) {
		binding["CurrentUser"] = currentUser(r)
	}

	renderer := render.New(render.Options{
		Layout: "layout",
	})
	renderer.HTML(w, http.StatusOK, templateName, binding)
}

func authorize(w http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, "auth")

	if session.Values["Email"] == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return false
	}

	return true
}

func currentUser(r *http.Request) models.User {
	session, _ := store.Get(r, "auth")
	user, _ := models.FindUser(session.Values["Email"].(string))
	return user
}

func loggedIn(r *http.Request) bool {
	session, _ := store.Get(r, "auth")
	return session.Values["Email"] != nil
}

func setSession(message interface{}, w http.ResponseWriter,
	r *http.Request, key ...string) {

	k := "Flash"

	if len(key) > 0 {
		k = key[0]
	}

	session, _ := store.Get(r, "auth")
	session.Values[k] = message
	session.Save(r, w)
}

func usernameConflictsWithRoute(username string) bool {
	for _, route := range routes() {
		if strings.Contains(route.Path, username) {
			return true
		}
	}

	return false
}
