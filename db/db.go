package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password123"
	dbname   = "postgres"
)

type FileStats struct {
	ID      int `json:"id"`
	Vowels  int `json:"vowels"`
	Capital int `json:"capital"`
	Small   int `json:"small"`
	Spaces  int `json:"spaces"`
}

var DbConn *sql.DB

func ConnectDB() error {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	DbConn, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	err = DbConn.Ping()
	if err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	fmt.Println("Connected to the database!")
	return nil
}
