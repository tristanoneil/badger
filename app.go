package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/joho/godotenv/autoload"
	"github.com/justinas/nosurf"
)

var (
	store *sessions.CookieStore
)

func main() {
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

	router := mux.NewRouter()
	router.HandleFunc("/", gists).Methods("GET")
	router.HandleFunc("/gists/new", newGist).Methods("GET")
	router.HandleFunc("/gists", newGist).Methods("POST")
	router.HandleFunc("/signup", signup).Methods("GET", "POST")
	router.HandleFunc("/login", login).Methods("GET", "POST")

	http.Handle("/", router)

	log.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nosurf.New(router))
}
