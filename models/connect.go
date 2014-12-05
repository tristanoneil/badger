package models

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var (
	db *sqlx.DB
)

func init() {
	db, _ = sqlx.Connect(
		"postgres",
		fmt.Sprintf("dbname=%s sslmode=disable", os.Getenv("DATABASE")),
	)
}
