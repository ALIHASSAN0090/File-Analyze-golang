package db

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteRecords(g *gin.Context) {

	idstr := g.PostForm("id")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value"})

		return
	}

	query := `DELETE FROM file_stats WHERE id = $1`

	_, err = DbConn.Exec(query, id)
	if err != nil {
		fmt.Println("error ", err)
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating record"})

		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Record deleted successfully"})

}
