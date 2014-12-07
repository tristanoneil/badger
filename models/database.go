package models

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var (
	db *sqlx.DB
)

func init() {
	var err error
	db, err = sqlx.Connect(
		"postgres",
		fmt.Sprintf("dbname=%s sslmode=disable", os.Getenv("DATABASE")),
	)

	if err != nil {
		log.Fatal(err)
	}

	MigrateDB()
}

func MigrateDB() {
	files, _ := ioutil.ReadDir("./migrations")

	for _, f := range files {
		_, err := sqlx.LoadFile(db, fmt.Sprintf("migrations/%s", f.Name()))

		if err != nil {
			log.Fatal(err)
		}
	}
}

func ResetDB() {
	db.MustExec("drop schema public cascade")
	db.MustExec("create schema public")
}
