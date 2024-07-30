package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectDB_Success(t *testing.T) {
	DbConn, err := Connect()
	defer DbConn.Close()
	assert.NoError(t, err, "Expected no error when connecting to the database")

	var result int
	err = DbConn.QueryRow("SELECT 1").Scan(&result)
	assert.NoError(t, err, "Expected no error when querying the database")

	fmt.Println("Database connection test passed!")
}

// func TestConnectDB_Failure(t *testing.T) {

// 	// DbConn = nil

// 	err := ConnectDB()
// 	assert.Error(t, err, "Expected error when failing to connect to the database")
// 	assert.Contains(t, err.Error(), "error opening database", "Expected specific error message")

// 	fmt.Println("Database connection failure test passed!")
// }
