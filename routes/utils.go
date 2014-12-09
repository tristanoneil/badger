package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/sessions"
	"github.com/justinas/nosurf"
	"github.com/tristanoneil/badger/models"
)

var (
	store *sessions.CookieStore
)

func init() {
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
}

func render(templateName string, w http.ResponseWriter,
	r *http.Request, binding map[string]interface{}) {

	session, _ := store.Get(r, "auth")

	if flash, ok := session.Values["Flash"].(string); ok {
		binding["Flash"] = flash
		setSession("", w, r)
	}

	binding["Token"] = nosurf.Token(r)

	if loggedIn(r) {
		binding["CurrentUserID"] = currentUserID(r)
	}

	templateBox, err := rice.FindBox("../templates")

	if err != nil {
		log.Fatal(err)
	}

	lstr, _ := templateBox.String("layout.tmpl")
	tstr, _ := templateBox.String(fmt.Sprintf("%s.tmpl", templateName))
	lstr += tstr

	var t *template.Template
	t, err = template.New("layout").Parse(lstr)

	if err != nil {
		log.Fatal(err)
	}

	err = t.ExecuteTemplate(w, "base", binding)

	if err != nil {
		log.Print("Error executing template: ", err)
	}
}

func authorize(w http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, "auth")

	if session.Values["Email"] == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return false
	}

	return true
}

func currentUserID(r *http.Request) int {
	session, _ := store.Get(r, "auth")
	return models.GetUserIDForEmail(session.Values["Email"].(string))
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
