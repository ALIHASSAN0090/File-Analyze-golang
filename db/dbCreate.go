package db

import "fmt"

func CreateUser(id, vowels, capital, small, spaces int) error {
	fmt.Println("In create user function")

	_, err := DbConn.Exec("INSERT INTO file_stats (user_id, vowels, capital, small, spaces) VALUES ($1, $2, $3, $4, $5)", id, vowels, capital, small, spaces)
	if err != nil {
		fmt.Println("Error executing query:", err) // Debug statement
		return fmt.Errorf("error inserting into database: %v", err)
	}

	return nil
}

func CreateUserData(name, password string) error {
	_, err := DbConn.Exec("INSERT INTO users (name,password) VALUES ($1, $2)", name, password)
	if err != nil {
		fmt.Println("Error executing query:", err) // Debug statement
		return fmt.Errorf("error inserting user into database: %v", err)
	}
	return nil
}
