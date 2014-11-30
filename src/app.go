package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	db *sqlx.DB
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {
	db, _ = sqlx.Connect("postgres", "dbname=badger sslmode=disable")

	router := mux.NewRouter()
	router.HandleFunc("/signup", signup).Methods("GET", "POST")
	router.HandleFunc("/login", login).Methods("GET", "POST")
	router.HandleFunc("/gists", gists).Methods("GET", "POST")

	http.Handle("/", router)

	log.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}

type Gist struct {
	Id      int    `db:"id"`
	UserId  int    `db:"user_id"`
	Title   string `db:"title"`
	Content string `db:"content"`
}

func gists(w http.ResponseWriter, r *http.Request) {
	authorize(w, r)
	gists := []Gist{}
	db.Select(&gists, "SELECT * FROM gists WHERE user_id = $1", currentUserId(r))
	render("gists.html", w, r, map[string]interface{}{"Gists": gists})
}
