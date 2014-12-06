package routes

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/justinas/nosurf"
	"github.com/tristanoneil/badger/models"
	r "github.com/unrolled/render"
)

var (
	store    *sessions.CookieStore
	renderer *r.Render
)

func init() {
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	renderer = r.New(r.Options{
		Layout: "layout",
	})
}

func render(template string, w http.ResponseWriter, r *http.Request, binding map[string]interface{}) {
	session, _ := store.Get(r, "auth")

	if flash, ok := session.Values["Flash"].(string); ok {
		binding["Flash"] = flash
		setSession("", w, r)
	}

	binding["Token"] = nosurf.Token(r)

	renderer.HTML(w, http.StatusOK, template, binding)
}

func authorize(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "auth")

	if session.Values["Email"] == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}

func currentUserID(r *http.Request) int {
	session, _ := store.Get(r, "auth")

	if email, ok := session.Values["Email"].(string); ok {
		return models.GetUserIDForEmail(email)
	}

	return 0
}

func setSession(message string, w http.ResponseWriter,
	r *http.Request, key ...string) {

	k := "Flash"

	if len(key) > 0 {
		k = key[0]
	}

	session, _ := store.Get(r, "auth")
	session.Values[k] = message
	session.Save(r, w)
}
