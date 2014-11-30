package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/justinas/nosurf"
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
	http.ListenAndServe(":3000", nosurf.New(router))
}
