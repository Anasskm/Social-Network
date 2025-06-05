package models

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"                // Replace with your MySQL username
	dbPass := os.Getenv("MY_SQLPW") // Replace with your MySQL password
	dbName := "new_schema"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(localhost:3306)/"+dbName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
