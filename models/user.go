package models

type User struct {
	ID                 string
	Name               string // unique within the database
	DisplayName        string
	Bio                string
	Type               string // user, organization, bot?
	Email              string
	Password           string
	AmountRepositories int
	// Status string // type of user
	// Picture string // profile picture full location URL
	// DefaultBranchName string // the name of the branch that all repos created by user will have
}
