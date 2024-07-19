package db

import "fmt"

func CreateUser(vowels, capital, small, spaces int) error {
	fmt.Println("In create user function")

	_, err := DbConn.Exec("INSERT INTO file_stats (vowels, capital, small, spaces) VALUES ($1, $2, $3, $4)", vowels, capital, small, spaces)
	if err != nil {
		return fmt.Errorf("error inserting into database: %v", err)
	}

	return nil
}
