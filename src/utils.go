package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func render(tmpl string, w http.ResponseWriter, r *http.Request, data ...map[string]interface{}) {
	tmpl = fmt.Sprintf("templates/%s", tmpl)
	t, err := template.ParseFiles(tmpl)

	if err != nil {
		log.Print("Template parsing error: ", err)
	}

	d := map[string]interface{}{}

	if len(data) > 0 {
		d = data[0]
	}

	session, _ := store.Get(r, "auth")

	if str, ok := session.Values["Flash"].(string); ok {
		d["Flash"] = str
	}

	err = t.Execute(w, d)

	if err != nil {
		log.Print("Template executing error: ", err)
	}
}

func authorize(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "auth")

	if session.Values["Email"] == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}

func currentUserId(r *http.Request) int {
	session, _ := store.Get(r, "auth")

	var Id int
	db.Get(&Id, "SELECT id FROM users WHERE email = $1", session.Values["Email"])

	return Id
}
