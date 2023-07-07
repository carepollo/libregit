package git

import (
	"io"
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

// get the content of the root README.md of given repo
func GetReadme(ownerName, repoName string) (string, error) {
	path := filepath.Join(utils.GlobalEnv.GitRoot, ownerName, repoName+".git")
	repo, err := git.PlainOpen(path)
	if err != nil {
		return "", err
	}

	ref, err := repo.Head()
	if err != nil {
		return "", err
	}

	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return "", err
	}

	tree, err := commit.Tree()
	if err != nil {
		return "", err
	}

	fileEntry, err := tree.FindEntry("README.md")
	if err != nil {
		return "", err
	}

	file, err := repo.BlobObject(fileEntry.Hash)
	if err != nil {
		return "", err
	}

	reader, err := file.Reader()
	if err != nil {
		return "", err
	}

	content, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// func AddReadme(ownerName, repoName string) error
