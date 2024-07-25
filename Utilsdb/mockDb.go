// testutils/mockdb.go
package testutils

import (
	"database/sql"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
)

func MockDB() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to create.. mock : %v", err)
	}

	return db, mock, nil
}
