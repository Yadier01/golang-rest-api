package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		log.Fatalf("could not connect to the db: %v", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTable()
}



func createTable() {
	createEventTable := `
 CREATE TABLE IF NOT EXISTS events (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL, 
  location TEXT NOT NULL,
  description TEXT NOT NULL,
  dateTime DATETIME NOT NULL,
  user_id INTEGER
  )
  `
	_, err := DB.Exec(createEventTable)
	if err != nil {
		log.Fatalf("could not create table: %v", err)
	}
}
