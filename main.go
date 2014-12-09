package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/justinas/nosurf"
	"github.com/tristanoneil/badger/routes"
)

func main() {
	n := negroni.Classic()
	n.UseHandler(nosurf.New(routes.Handlers()))
	log.Println(fmt.Sprintf("Listening on port %s", os.Getenv("PORT")))
	n.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
