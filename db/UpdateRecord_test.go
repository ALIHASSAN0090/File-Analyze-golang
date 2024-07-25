package db

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUpdateRecord_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("id=1&value=5"))
	ctx.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	DbConn = db

	mock.ExpectExec(`UPDATE file_stats SET vowels = \$1 WHERE id = \$2`).
		WithArgs(5, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	UpdateRecord(ctx)

	assert.Equal(t, http.StatusOK, w.Code, "HTTP status code should be OK")
	expectedResponse := `{"message":"Record updated successfully"}`
	assert.JSONEq(t, expectedResponse, w.Body.String(), "Response body should match expected JSON")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unmet expectations: %s", err)
	}
}

func TestUpdateRecord_Failure_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("id=invalid&value=5"))
	ctx.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	UpdateRecord(ctx)

	assert.Equal(t, http.StatusBadRequest, w.Code, "HTTP status code should be BadRequest")
	expectedResponse := `{"error":"Invalid ID"}`
	assert.JSONEq(t, expectedResponse, w.Body.String(), "Response body should match expected JSON")
}

func TestUpdateRecord_Failure_InvalidValue(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("id=1&value=invalid"))
	ctx.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	UpdateRecord(ctx)

	assert.Equal(t, http.StatusBadRequest, w.Code, "HTTP status code should be BadRequest")
	expectedResponse := `{"error":"Invalid value"}`
	assert.JSONEq(t, expectedResponse, w.Body.String(), "Response body should match expected JSON")
}

func TestUpdateRecord_Failure_DBError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("id=1&value=5"))
	ctx.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	DbConn = db

	mock.ExpectExec(`UPDATE file_stats SET vowels = \$1 WHERE id = \$2`).
		WithArgs(5, 1).
		WillReturnError(fmt.Errorf("some database error"))

	UpdateRecord(ctx)

	assert.Equal(t, http.StatusInternalServerError, w.Code, "HTTP status code should be InternalServerError")
	expectedResponse := `{"error":"Error updating record"}`
	assert.JSONEq(t, expectedResponse, w.Body.String(), "Response body should match expected JSON")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unmet expectations: %s", err)
	}
}
