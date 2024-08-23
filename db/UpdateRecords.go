package db

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateRecord godoc
// @Summary Update a file statistics record
// @Description Update the vowel count of a file statistics record by ID.
// @Accept  multipart/form-data
// @Produce json
// @Param id formData int true "ID of the record to update"
// @Param value formData int true "New vowel count"
// @Success 200 {object} map[string]string "Record updated successfully"
// @Failure 400 {object} map[string]string "Invalid ID or value"
// @Failure 500 {object} map[string]string "Error updating record"
// @Router / [put]
func UpdateRecord(g *gin.Context) {
	var updateData struct {
		ID    int `json:"id"`
		Value int `json:"value"`
	}

	// Log the incoming request
	fmt.Println("Incoming Request Body:", g.Request.Body)

	if err := g.ShouldBindJSON(&updateData); err != nil {
		fmt.Println("Binding Error:", err) // Log binding errors
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	fmt.Println("Update Data:", updateData) // Log the data being used in the query

	sqlStatement := `UPDATE file_stats SET vowels = $1 WHERE id = $2`

	_, err := DbConn.Exec(sqlStatement, updateData.Value, updateData.ID)
	if err != nil {
		fmt.Println("Database Error:", err) // Log database errors
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating record"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Record updated successfully"})
}
