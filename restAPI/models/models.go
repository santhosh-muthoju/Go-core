package models

import (
	"fmt"
	"interviewPrep/restAPI/db"
	"log"
	"time"
)

type Event struct {
	Id          int64     `json:"ID"`
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	Location    string    `json:"Location"`
	DateTime    time.Time `json:"DateTime"`
	UserId      int       `json:"UserId"`
}

func Save(event Event) (int64, error) {
	// ✅ Ensure DateTime is properly formatted
	var formattedDateTime string
	if event.DateTime.IsZero() {
		formattedDateTime = time.Now().Format("2006-01-02 15:04:05")
	} else {
		formattedDateTime = event.DateTime.Format("2006-01-02 15:04:05")
	}

	// ✅ Ensure user_id is set correctly
	if event.UserId == 0 {
		event.UserId = 1 // Default to 1 (or another appropriate value)
	}

	query := `
		INSERT INTO events (name, description, location, dateTime, user_id)
		VALUES (?, ?, ?, ?, ?)
	`

	result, err := db.DB.Exec(query, event.Name, event.Description, event.Location, formattedDateTime, event.UserId)
	if err != nil {
		log.Printf("MySQL Insert Error: %v", err)
		return 0, fmt.Errorf("failed to insert event: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error fetching last insert ID: %v", err)
		return 0, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	event.Id = lastInsertID
	log.Printf("Event saved successfully: %+v\n", event)
	return lastInsertID, nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT id, name, description, location, dateTime, user_id FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		log.Println("MySQL Query Error:", err)
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		var dateTimeStr string // Store as string to parse later

		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &dateTimeStr, &event.UserId)
		if err != nil {
			log.Println("Row Scan Error:", err)
			return nil, err
		}

		// ✅ Convert MySQL DATETIME string to time.Time
		event.DateTime, err = time.Parse("2006-01-02 15:04:05", dateTimeStr)
		//or also use this event.DateTime, err = time.Parse(time.RFC3339, dateTimeStr)
		if err != nil {
			log.Println("DateTime Parse Error:", err)
			event.DateTime = time.Now() // Default to current time if parsing fails
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		log.Println("Rows error:", err)
		return nil, err
	}

	return events, nil
}
