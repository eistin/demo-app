package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Open database connection
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DBNAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("Error connecting to MySQL:", err)
	}

	// Check if the database connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging MySQL database:", err)
	}

	defer db.Close()

	log.Println("DB Connection OK")

	// Routes
	http.HandleFunc("/increment", incrementHandler)
	http.HandleFunc("/decrement", decrementHandler)
	http.HandleFunc("/current", currentHandler)

	// Start server
	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func incrementHandler(w http.ResponseWriter, r *http.Request) {
	// Increment the counter in the database
	log.Println("Incr")
	_, err := db.Exec("UPDATE counter SET count = count + 1")
	if err != nil {
		log.Println("Error incrementing counter:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(w, "Counter incremented")
}

func decrementHandler(w http.ResponseWriter, r *http.Request) {
	// Decrement the counter in the database
	log.Println("Decr")
	_, err := db.Exec("UPDATE counter SET count = count - 1")
	if err != nil {
		log.Println("Error decrementing counter:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(w, "Counter decremented")
}

func currentHandler(w http.ResponseWriter, r *http.Request) {
	// Get the current counter value from the database
	var count int
	log.Println("Current")
	err := db.QueryRow("SELECT count FROM counter").Scan(&count)
	if err != nil {
		log.Println("Error getting current counter:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write the counter value as the response
	fmt.Fprintf(w, "%d", count)
}

// addCorsHeaders adds the necessary CORS headers to the HTTP handler.
// func addCorsHeaders(handler http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

// 		// Call the next handler in the chain
// 		handler.ServeHTTP(w, r)
// 	})
// }
