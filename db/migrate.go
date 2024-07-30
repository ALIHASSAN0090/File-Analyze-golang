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
	query := `CREATE TABLE IF NOT EXISTS file_stats (
        id SERIAL PRIMARY KEY,
        vowels INT,
        capital INT,
        small INT,
        spaces INT
    )`

	_, err := DbConn.Exec(query)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error Creating table"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"error": "Created Table succesfully"})
}
