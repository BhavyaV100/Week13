package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Connect to MySQL database
	setupDB()

	// Set up Gin router
	router := gin.Default()

	// Define API endpoint
	router.GET("/current-time", getCurrentTime)

	// Run the server
	router.Run(":9091")
}

func setupDB() {
	// Replace 'user:password@tcp(127.0.0.1:3306)/toronto_time_db' with your MySQL connection details
	dsn := "root:root@tcp(127.0.0.1:3306)/goapi"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")
}

func getCurrentTime(c *gin.Context) {
	// Get current time in Toronto's timezone
	torontoTime, err := timeInToronto()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Toronto time"})
		return
	}

	// Insert current time into the database
	_, err = db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", torontoTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log time"})
		return
	}

	// Return current time in JSON format
	c.JSON(http.StatusOK, gin.H{"current_time": torontoTime})
}

func timeInToronto() (time.Time, error) {
	loc, err := time.LoadLocation("America/Toronto")
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().In(loc), nil
}
