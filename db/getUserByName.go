package db

import "main.go/models"

func GetUserByName(name string) (*models.User, error) {
	query := `SELECT id, name, password FROM users WHERE name = $1`
	row := DbConn.QueryRow(query, name)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
