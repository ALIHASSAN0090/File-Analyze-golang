package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	testutils "main.go/Utilsdb"
	"main.go/db"
)

func TestGetRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)

	MockDB, mock, err := testutils.MockDB()
	if err != nil {
		t.Fatalf("Error in connection of mock database: %v", err)
	}
	defer MockDB.Close()

	db.DbConn = MockDB

	expectedRows := sqlmock.NewRows([]string{"id", "vowels", "capital", "small", "spaces"}).
		AddRow(1, 3, 5, 7, 2).
		AddRow(5, 4, 3, 6, 1)

	mock.ExpectQuery("SELECT \\* FROM file_stats").WillReturnRows(expectedRows)

	r := gin.Default()
	r.GET("/", DisplayAll)

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

	expectedBody := `[{"id":1,"vowels":3,"capital":5,"small":7,"spaces":2},{"id":5,"vowels":4,"capital":3,"small":6,"spaces":1}]`
	assert.JSONEq(t, expectedBody, response.Body.String())

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unmet mock expectations: %v", err)
	}
}
