package db

import (
	"fmt"

	"github.com/carepollo/librecode/models"
	"github.com/carepollo/librecode/utils"
	"github.com/google/uuid"
)

func GetReposByOwner(ownerId string) ([]models.Repo, error) {
	repos := []models.Repo{}
	rows, err := db.Query("SELECT * FROM repos WHERE owner = ?", ownerId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		repo := models.Repo{}
		err := rows.Scan(
			&repo.ID,
			&repo.Name,
			&repo.Description,
			&repo.IsPublic,
			&repo.Owner,
		)
		if err != nil {
			return nil, err
		}

		repos = append(repos, repo)
	}

	return repos, nil
}

func CreateRepo(userId, name string, visibility bool) error {
	id := uuid.New().String()
	location := fmt.Sprintf("%v/%v/%v", utils.GlobalEnv.URLs.Project, userId, name)
	query := "INSERT INTO repos (name, owner, isPublic, id, location) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, name, userId, visibility, id, location)
	return err
}
