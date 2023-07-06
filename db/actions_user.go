package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/carepollo/librecode/models"
)

var (
	// use when storing or retrieven data from cache about user being logged in or not
	sessionPrefix = "session_"
)

// helper to return empty value of user if error and assign properties from result to go struct
func extractSingleUser(result sqlResult) (models.User, error) {
	user := models.User{}
	err := result.Scan(
		&user.ID,
		&user.Name,
		&user.DisplayName,
		&user.Bio,
		&user.Type,
		&user.Email,
		&user.Password,
		&user.AmountRepositories,
		&user.Status,
		&user.Picture,
		&user.DefaultBranchName,
	)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// use for register, it creates a new user in DB using assumed values
func CreateUser(id, name, email, password, picture string) error {
	query := "INSERT INTO users (id, name, email, password, picture) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, id, name, email, password, picture)

	return err
}

// use for login, give as param a value that could be username or email of the searched user
func GetUserByNameOrEmail(userOrMail string) (models.User, error) {
	query := "SELECT * FROM users WHERE name = ? OR email = ?"
	result := db.QueryRow(query, userOrMail, userOrMail)
	user, err := extractSingleUser(result)
	if err != nil {
		log.Printf("%v", err.Error())
		return models.User{}, err
	}

	return user, nil
}

// abstraction to search the existance of a user by given values, use internally only
func userExistsBy(param string, value string) bool {
	user := models.User{}
	query := fmt.Sprintf("SELECT id FROM users WHERE %v = ?", param)
	result := db.QueryRow(query, value)
	err := result.Scan(&user.ID)

	return err == nil
}

// check if a username exists
func NameIsRegistered(name string) bool {
	return userExistsBy("name", name)
}

// check if a user with given email exists
func EmailIsRegistered(email string) bool {
	return userExistsBy("email", email)
}

// activate user, use when verify
func ActivateUser(id string) error {
	query := "UPDATE users SET status = ? WHERE id = ?"
	result, err := db.Exec(query, models.ACTIVE, id)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("rows not affected")
	}

	return err
}

// function to use on login. Is to store the user data to persist across session
func SetUserSession(deviceId string, user models.User) error {
	duration := time.Hour * 24
	session := sessionPrefix + deviceId
	return remember(session, user, duration)
}

// function to use on any logged-in route, to get context of the logged user data.
func GetUserSession(deviceId string) (models.User, error) {
	result, err := retrieve(sessionPrefix + deviceId)
	if err != nil {
		return models.User{}, err
	}

	// unencoding from json to go struct
	user := models.User{}
	err = json.Unmarshal(result, &user)
	if err != nil {
		log.Println(err)
		return models.User{}, err
	}
	return user, nil
}

func DeleteUserSession(deviceId string) error {
	return forget(sessionPrefix + deviceId)
}

// use to retrieve data of user
func GetUserByName(username string) (models.User, error) {
	query := "SELECT * FROM users WHERE name = ?"
	result := db.QueryRow(query, username)
	user, err := extractSingleUser(result)
	return user, err
}

// replace picture, bio, displayname, amountrepositores, and password with given values
func UpdateUser(data models.User) error {
	query := "UPDATE users SET picture = ?, bio = ?, displayName = ?, amountRepositories = ?, password = ? WHERE id = ?"
	_, err := db.Exec(
		query,
		data.Picture,
		data.Bio,
		data.DisplayName,
		data.AmountRepositories,
		data.Password,
		data.ID,
	)

	return err
}
