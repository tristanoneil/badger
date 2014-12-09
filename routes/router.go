package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", gists).Methods("GET")
	router.HandleFunc("/gists/new", newGist).Methods("GET")
	router.HandleFunc("/gists", newGist).Methods("POST")
	router.HandleFunc("/signup", signup).Methods("GET", "POST")
	router.HandleFunc("/login", login).Methods("GET", "POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))
	return router
}
