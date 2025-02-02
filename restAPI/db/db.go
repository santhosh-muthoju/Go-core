package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	dsn := "root:RootRoot@tcp(localhost:3306)/events_db"

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	fmt.Println("MySQL database connected successfully.")
	createTables()
}

func createTables() {
	query := `
	CREATE TABLE IF NOT EXISTS events (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		location VARCHAR(255) NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INT
	);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err) // Proper error logging
	}
	fmt.Println("Tables checked/created successfully.")
}
