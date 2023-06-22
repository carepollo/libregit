package git

import (
	"os"

	"github.com/go-git/go-git/v5"
)

// create a directory in given path and initialize a server-side repository
// in the git root location
func CreateRepo(path string) (*git.Repository, error) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return nil, err
	}

	return git.PlainInit(path, true)
}
