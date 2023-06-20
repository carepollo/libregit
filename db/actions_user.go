package db

import (
	"log"

	"github.com/carepollo/librecode/models"
	"github.com/google/uuid"
)

// use for register
func CreateUser(name, email, password string) error {
	query := "INSERT INTO users (id, name, amountRepositories, type, email, password) VALUES (?, ?, 1, 'user', ?, ?)"
	id := uuid.New().String()
	_, err := client.Exec(query, id, name, email, password)

	return err
}

// use for login
func GetUserByNameOrEmail(entry string) *models.User {
	user := models.User{}
	result := client.QueryRow("SELECT * FROM users WHERE name = ? OR email = ?", entry, entry)
	err := result.Scan(
		&user.ID,
		&user.Name,
		&user.DisplayName,
		&user.Bio,
		&user.Type,
		&user.Email,
		&user.Password,
		&user.AmountRepositories,
	)

	if err != nil {
		log.Printf("%v", err.Error())
		return nil
	}

	return &user
}
