package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	client *sql.DB
)

// start connection with mysql database (main)
func Open(connection string) {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic("could not connect to database")
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping: %v", err)
	}

	client = db
	log.Println("Successfully connected to MYSQL instance")
}

// MUST be executed before program exits and after each opening of database as a DEFER
func Close() {
	client.Close()
}
