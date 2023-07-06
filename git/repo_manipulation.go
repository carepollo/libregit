package git

import (
	"os"
	"path/filepath"

	"github.com/carepollo/librecode/utils"
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

// rename the folder of a certain repo by given path
func RenameRepo(oldName, newName, owner string) error {
	oldPath := filepath.Join(utils.GlobalEnv.GitRoot, owner, oldName+".git")
	newPath := filepath.Join(utils.GlobalEnv.GitRoot, owner, newName+".git")

	if err := os.Rename(oldPath, newPath); err != nil {
		return err
	}
	return nil
}
