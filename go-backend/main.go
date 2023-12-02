package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {
	// Open the database connection
	db, err := sql.Open("postgres", "user=root dbname=QR sslmode=disable password=qareebpassword host=localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Load test data (Optional if you've already loaded data into the database)
	// loadTestData(db)

	// Set up chi router
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Define your routes
	router.Post("/update_qr_code/{code}", func(w http.ResponseWriter, r *http.Request) {
		qrCode := chi.URLParam(r, "code")
		updateQRCodeState(db, w, r, qrCode)
	})

	// Start the HTTP server
	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func updateQRCodeState(db *sql.DB, w http.ResponseWriter, r *http.Request, qrCode string) {
	// Update the database
	_, err := db.Exec("UPDATE qr_codes SET is_active = false WHERE code = $1", qrCode)
	if err != nil {
		log.Println("Error updating QR code state:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
