package models

type Repo struct {
	ID          string // unique id within the entire application
	Name        string // unique within the owner's scope
	Description string // description of repo (non git feature)
	Owner       string // owner of the repo, who created it, contains its ID
	Location    string // full site location URL within the system "username/reponame"
	IsPrivate   bool   // to handle read/write permissions
	// LastUpdate time.Time // last date the repo a change was detected in the repo.
}
