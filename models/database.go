package models

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/jmoiron/sqlx"
	"github.com/tristanoneil/badger/static"

	//
	// Allows sqlx to connect to Postgres.
	//
	_ "github.com/lib/pq"
)

var (
	Db *sqlx.DB
)

func init() {
	var err error
	Db, err = sqlx.Connect(
		"postgres",
		fmt.Sprintf("dbname=%s sslmode=disable user=%s password=%s",
			os.Getenv("DATABASE"),
			os.Getenv("DATABASE_USER"),
			os.Getenv("DATABASE_PASSWORD"),
		))

	if err != nil {
		log.Fatal(err)
	}

	MigrateDB()
}

//
// MigrateDB loads all SQL files from migrations and executes them.
//
func MigrateDB() {
	for _, f := range static.AssetNames() {
		match, _ := regexp.MatchString("migrations", f)

		if match {
			sql, _ := static.Asset(f)
			Db.MustExec(string(sql))
		}
	}
}

//
// ResetDB resets the database schema, useful for testing.
//
func ResetDB() {
	Db.MustExec("drop schema public cascade")
	Db.MustExec("create schema public")
}
