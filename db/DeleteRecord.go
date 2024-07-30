package db

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteRecords godoc
// @Summary Delete a file statistics record
// @Description Delete a file statistics record by ID.
// @Accept  multipart/form-data
// @Produce json
// @Param id formData int true "ID of the record to delete"
// @Success 200 {object} map[string]string "Record deleted successfully"
// @Failure 400 {object} map[string]string "Invalid ID"
// @Failure 500 {object} map[string]string "Error deleting record"
// @Router /delete [Delete]
func DeleteRecords(g *gin.Context) {
	idStr := g.PostForm("id")

	id, err := strconv.Atoi(idStr)
	fmt.Print(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value"})
		return
	}

	query := "DELETE FROM file_stats WHERE id = $1"
	_, err = DbConn.Exec(query, id)
	if err != nil {
		fmt.Println("error ", err)
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting record"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Record deleted successfully"})
}
