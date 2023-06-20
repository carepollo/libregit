package models

type Repo struct {
	ID          int    // unique id within the entire application
	Name        string // unique within the owner's scope
	Description string // description of repo (non git feature)
	IsPublic    bool   // to handle read/write permissions
	Owner       int    // owner of the repo, who created it, contains its ID
}
