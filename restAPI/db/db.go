// package db

// import (
// 	"database/sql"

// 	_ "github.com/mattn/go-sqlite3"
// )

// var DB *sql.DB

// func InitDB() {
// 	var err error
// 	DB, err = sql.Open("sqlite3", "api.db")

// 	if err != nil {
// 		panic("Could not connect to database.")
// 	}

// 	DB.SetMaxOpenConns(10)
// 	DB.SetMaxIdleConns(5)

// 	createTables()
// }

// func createTables() {
// 	createEventsTable := `
// 	CREATE TABLE IF NOT EXISTS events (
// 	  id INTEGER PRIMARY KEY AUTOINCREMENT,
// 	  name TEXT NOT NULL,
// 	  description TEXT NOT NULL,
// 	  location TEXT NOT NULL,
// 	  dateTime DATETIME NOT NULL,
// 	  user_id INTEGER
// 	)
// 	`

// 	_, err := DB.Exec(createEventsTable)

// 	if err != nil {
// 		panic("could not create the events table.")
// 	}
// }

package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite3 driver

	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") // Ensure "api.db" is accessible

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err) // Show actual error
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	  id INTEGER PRIMARY KEY AUTOINCREMENT,
	  name TEXT NOT NULL,
	  description TEXT NOT NULL,
	  location TEXT NOT NULL,
	  dateTime DATETIME NOT NULL,
	  user_id INTEGER 
	);` // Added ';' at the end (good practice)

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		log.Fatalf("Error creating events table: %v", err) // Print actual error
	}
}
