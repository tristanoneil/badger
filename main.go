package main

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/tristanoneil/badger/routes"
)

func main() {
	log.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nosurf.New(routes.Handlers()))
}
