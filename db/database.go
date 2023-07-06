package db

import (
	"database/sql"
	"log"

	"github.com/carepollo/librecode/utils"
	_ "github.com/go-sql-driver/mysql"
)

// should satisfy *sql.Row and *sql.Rows structs
type sqlResult interface {
	Scan(dest ...interface{}) error
}

var db *sql.DB

// start connection with mysql database (main)
func openDatabase() {
	client, err := sql.Open("mysql", utils.GlobalEnv.Storage.Db.Connection)
	if err != nil {
		panic("could not connect to database")
	}

	if err := client.Ping(); err != nil {
		log.Fatalf("failed to ping: %v", err)
	}

	db = client
	log.Println("Successfully connected to MySQL instance")
}

// MUST be executed before program exits and after each opening of database as a DEFER
func closeDatabase() {
	if err := db.Close(); err != nil {
		log.Println(err)
	}
}
