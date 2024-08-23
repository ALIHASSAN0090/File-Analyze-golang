package db

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTable godoc
// @Summary Create the file statistics table
// @Description Create the file_stats table if it does not exist.
// @Produce json
// @Success 200 {object} map[string]string "Created Table successfully"
// @Failure 500 {object} map[string]string "Error Creating table"
// @Router /create-table [post]
func Createtable(g *gin.Context) {
	queryFile := `CREATE TABLE IF NOT EXISTS file_stats (
        id SERIAL PRIMARY KEY,
        vowels INT,
        capital INT,
        small INT,
        spaces INT
    )`

	_, err := DbConn.Exec(queryFile)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error Creating table for File Stats"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"Success": "Created Table succesfully for File Stats"})
}

func CreateTableUser(g *gin.Context) {
	// Define the SQL query with corrected data types and no trailing comma
	queryUser := `CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        password VARCHAR(255) NOT NULL
    )`

	// Execute the SQL query
	_, err := DbConn.Exec(queryUser)
	if err != nil {
		// Respond with an error if the query fails
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error Creating table for Users"})
		return
	}

	// Respond with success if the query succeeds
	g.JSON(http.StatusOK, gin.H{"Success": "Created Table successfully for Users"})
}
