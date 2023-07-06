package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/carepollo/librecode/models"
	"github.com/carepollo/librecode/utils"
	"github.com/google/uuid"
)

func extractSingleRepo(result sqlResult) (models.Repo, error) {
	repo := models.Repo{}
	err := result.Scan(
		&repo.ID,
		&repo.Name,
		&repo.Description,
		&repo.IsPrivate,
		&repo.Owner,
		&repo.Location,
	)
	if err != nil {
		return models.Repo{}, err
	}

	return repo, nil
}

// get all repos who has been created by given username
func GetReposByOwner(ownerName string, includePrivate bool) ([]models.Repo, error) {
	var err error
	var rows *sql.Rows
	repos := []models.Repo{}
	query := "SELECT * FROM repos WHERE owner = ?"

	if includePrivate {
		rows, err = db.Query(query, ownerName)
	} else {
		query += " AND isPrivate = ?"
		rows, err = db.Query(query, ownerName, includePrivate)
	}

	if err == sql.ErrNoRows {
		return []models.Repo{}, nil
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		repo, err := extractSingleRepo(rows)
		if err != nil {
			return nil, err
		}

		repos = append(repos, repo)
	}

	return repos, nil
}

// register on database a repo with given author, name and visibility
func CreateRepo(ownerName, repoName string, visibility bool) error {
	id := uuid.New().String()
	location := fmt.Sprintf("%v/%v/%v", utils.GlobalEnv.URLs.Project, ownerName, repoName)
	_, existsErr := GetRepoByOwnerAndName(ownerName, repoName)
	if existsErr == nil {
		errMessage := fmt.Sprintf("the repo %s/%s already exists", ownerName, repoName)
		return errors.New(errMessage)
	}

	query := "INSERT INTO repos (name, owner, isPrivate, id, location) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, repoName, ownerName, visibility, id, location)
	return err
}

// owner username and repo name to retrieve a unique repo
func GetRepoByOwnerAndName(owner, name string) (models.Repo, error) {
	result := db.QueryRow("SELECT * FROM repos WHERE owner = ? AND name = ?", owner, name)
	repo, err := extractSingleRepo(result)
	return repo, err
}

// update repo data in db, it only replaces
func UpdateRepo(data models.Repo) error {
	query := "UPDATE repos SET location = ?, name = ?, isPrivate = ?, description = ?"
	_, err := db.Exec(query, data.Location, data.Name, data.IsPrivate, data.Description)
	return err
}
