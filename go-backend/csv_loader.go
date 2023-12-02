// csv_loader.go

package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func loadTestData(db *sql.DB) {
	// Open the CSV file
	file, err := os.Open("example.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Parse the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through records and insert into the database
	for i, record := range records {
		// Skip the header row
		if i == 0 {
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Error converting 'id' to integer: %v\n", err)
			continue
		}

		code := record[1]
		isActive := record[2]

		// Insert into the qr_codes table
		_, err = db.Exec("INSERT INTO qr_codes (id, code, is_active) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING", id, code, isActive)
		if err != nil {
			log.Printf("Error inserting record: %v\n", err)
		}
	}

	fmt.Println("Test data loaded successfully.")
}
