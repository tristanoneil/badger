package routes

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

type route struct {
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
	Methods []string
}

func Router() *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes() {
		for _, method := range route.Methods {
			router.HandleFunc(route.Path, route.Handler).Methods(method)
		}
	}

	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("../assets").HTTPBox()))
	return router
}

func routes() []route {
	return []route{
		{Path: "/", Handler: gists, Methods: []string{"GET"}},
		{Path: "/gists/new", Handler: newGist, Methods: []string{"GET"}},
		{Path: "/gists", Handler: newGist, Methods: []string{"POST"}},
		{Path: "/gists/{id:[0-9]+}", Handler: showGist, Methods: []string{"GET"}},
		{Path: "/gists/{id:[0-9]+}", Handler: editGist, Methods: []string{"POST"}},
		{Path: "/gists/{id:[0-9]+}/edit", Handler: editGist, Methods: []string{"GET"}},
		{Path: "/gists/{id:[0-9]+}/delete", Handler: deleteGist, Methods: []string{"GET"}},
		{Path: "/signup", Handler: signup, Methods: []string{"GET", "POST"}},
		{Path: "/login", Handler: login, Methods: []string{"GET", "POST"}},
		{Path: "/logout", Handler: logout, Methods: []string{"GET"}},
		{Path: "/{username:[a-z]+}", Handler: usersGists, Methods: []string{"GET"}},
	}
}
