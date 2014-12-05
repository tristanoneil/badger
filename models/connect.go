package models

import (
	"fmt"
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
}
