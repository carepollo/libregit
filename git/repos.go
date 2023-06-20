package git

import "github.com/go-git/go-git/v5"

// create a server-side repository in the git root location
func CreateRepo(path string) (*git.Repository, error) {
	return git.PlainInit(path, true)
}
