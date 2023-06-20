package db

import "github.com/carepollo/librecode/models"

func GetReposByOwner(ownerId string) ([]models.Repo, error) {
	repos := []models.Repo{}
	rows, err := client.Query("SELECT * FROM repos WHERE owner = ?", ownerId)
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
