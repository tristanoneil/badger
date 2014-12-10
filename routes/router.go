package routes

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", gists).Methods("GET")
	router.HandleFunc("/gists/new", newGist).Methods("GET")
	router.HandleFunc("/gists", newGist).Methods("POST")
	router.HandleFunc("/gists/{id:[0-9]+}", showGist).Methods("GET")
	router.HandleFunc("/signup", signup).Methods("GET", "POST")
	router.HandleFunc("/login", login).Methods("GET", "POST")
	router.HandleFunc("/logout", logout).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("../assets").HTTPBox()))
	return router
}
