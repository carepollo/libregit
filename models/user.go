package models

type UserStatus string

const (
	// status of user when account has been created but not verified,
	// it cannnot login in this state.
	UNVERIFIED UserStatus = "unverified"
	// representation of the status of a user when is an activated account.
	// In this state it can do any regular user action.
	ACTIVE UserStatus = "active"
)

type User struct {
	ID                 string     `json:"id"`                 // unique ID within the system
	Name               string     `json:"name"`               // unique within the database
	DisplayName        string     `json:"displayName"`        // a non-unique alias for the user
	Bio                string     `json:"bio:"`               // user description, defaults to ''
	Type               string     `json:"type"`               // user, organization, bot. defaults to user
	Email              string     `json:"email"`              // unique email within the system
	Password           string     `json:"password"`           // hashed and salted password
	AmountRepositories int        `json:"amountRepositories"` // amount of repositories the user has, defaults to 1
	Status             UserStatus `json:"status"`             // type of user
	Picture            string     `json:"picture"`            // profile picture full location URL
	DefaultBranchName  string     `json:"defaultBranchName"`  // the name of the branch that all repos created by user will have
}
