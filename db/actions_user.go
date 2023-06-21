package db

import (
	"fmt"
	"log"

	"github.com/carepollo/librecode/models"
	"github.com/google/uuid"
)

// use for register, it creates a new user in DB using assumed values
func CreateUser(name, email, password, picture string) error {
	query := "INSERT INTO users (id, name, email, password, picture) VALUES (?, ?, ?, ?, ?)"
	id := uuid.New().String()
	_, err := client.Exec(query, id, name, email, password, picture)

	return err
}

// use for login
func GetUserByNameOrEmail(entry string) (models.User, error) {
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
		return models.User{}, err
	}

	return user, nil
}

// check if a username exists
func NameIsRegistered(name string) bool {
	return userExistsBy("name", name)
}

// check if a user with given email exists
func EmailIsRegistered(email string) bool {
	return userExistsBy("email", email)
}

// abstraction for user existing
func userExistsBy(param string, value string) bool {
	user := models.User{}
	query := fmt.Sprintf("SELECT id FROM users WHERE %v = ?", param)
	result := client.QueryRow(query, value)
	err := result.Scan(&user.ID)

	return err == nil
}
