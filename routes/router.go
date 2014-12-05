package routes

import "github.com/gorilla/mux"

func Handlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", gists).Methods("GET")
	router.HandleFunc("/gists/new", newGist).Methods("GET")
	router.HandleFunc("/gists", newGist).Methods("POST")
	router.HandleFunc("/signup", signup).Methods("GET", "POST")
	router.HandleFunc("/login", login).Methods("GET", "POST")
	return router
}
