package models

type UserStatus string

const (
	UNVERIFIED UserStatus = "unverified"
	ACTIVE     UserStatus = "active"
)

type User struct {
	ID                 string     // unique ID within the system
	Name               string     // unique within the database
	DisplayName        string     // a non-unique alias for the user
	Bio                string     // user description, defaults to ''
	Type               string     // user, organization, bot. defaults to user
	Email              string     // unique email within the system
	Password           string     // hashed and salted password
	AmountRepositories int        // amount of repositories the user has, defaults to 1
	Status             UserStatus // type of user
	Picture            string     // profile picture full location URL
	DefaultBranchName  string     // the name of the branch that all repos created by user will have
	// CreationDate time.DateTime // date when the user created account
}
