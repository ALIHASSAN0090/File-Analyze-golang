package db

import (
	"fmt"
	"net/http"
	"strconv"

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
	idStr := g.PostForm("id")
	valueStr := g.PostForm("value")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value"})
		return
	}

	sqlStatement := `UPDATE file_stats SET vowels = $1 WHERE id = $2`

	_, err = DbConn.Exec(sqlStatement, value, id)
	if err != nil {
		fmt.Println("error ", err)
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating record"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Record updated successfully"})
}
