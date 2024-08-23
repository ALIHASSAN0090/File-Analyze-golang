package db

import (
	"fmt"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestInsertValue_Success(t *testing.T) {
	// Mock DB connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	DbConn = db
	mock.ExpectExec("INSERT INTO file_stats").
		WithArgs(3, 5, 7, 2).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = CreateUser(3, 5, 7, 2, 2)

	assert.NoError(t, err, "CreateUser should not return an error")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unmet expectations: %s", err)
	}
}

func TestInsertValue_Failure(t *testing.T) {
	// Mock DB connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	DbConn = db

	mock.ExpectExec("INSERT INTO file_stats").
		WillReturnError(fmt.Errorf("error inserting into database"))

	err = CreateUser(3, 5, 7, 2, 2)

	assert.Error(t, err, "CreateUser should return an error")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unmet expectations: %s", err)
	}
}
