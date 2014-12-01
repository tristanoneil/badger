package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	"github.com/justinas/nosurf"
	_ "github.com/lib/pq"
)

var (
	db    *sqlx.DB
	store *sessions.CookieStore
)

func main() {
	db, _ = sqlx.Connect(
		"postgres",
		fmt.Sprintf("dbname=%s sslmode=disable", os.Getenv("DATABASE")),
	)

	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

	router := mux.NewRouter()
	router.HandleFunc("/", gists).Methods("GET", "POST")
	router.HandleFunc("/signup", signup).Methods("GET", "POST")
	router.HandleFunc("/login", login).Methods("GET", "POST")

	http.Handle("/", router)

	log.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nosurf.New(router))
}
