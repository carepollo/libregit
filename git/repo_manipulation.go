package git

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/carepollo/librecode/utils"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/filemode"
)

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

func GetRepoFile(ownerName, repoName, entryPath string) (string, error) {
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

	fileEntry, err := tree.FindEntry(entryPath)
	if err != nil {
		return "", err
	}

	if fileEntry.Mode != filemode.Regular {
		return "", errors.New("entry is not file")
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

// Get the files and directories present at given path, pass "" when expecting
// something locates at the root directory.
func GetRepoDir(ownername, reponame, path string) ([]map[string]interface{}, error) {
	result := []map[string]interface{}{}
	repodir := reponame + ".git"
	repopath := filepath.Join(utils.GlobalEnv.GitRoot, ownername, repodir)
	repo, err := git.PlainOpen(repopath)
	if err != nil {
		return nil, err
	}

	ref, err := repo.Head()
	if err != nil {
		return nil, err
	}

	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}

	tree, err := commit.Tree()
	if err != nil {
		return nil, err
	}

	if path != "" {
		tree, err = tree.Tree(path)
		if err != nil {
			return nil, err
		}
	}

	for _, entry := range tree.Entries {
		data := make(map[string]interface{})
		data["IsDirectory"] = entry.Mode.String() == filemode.Dir.String()
		data["Name"] = entry.Name
		data["Location"] = fmt.Sprintf(
			"%s/%s/%s/%s/%s/%s",
			utils.GlobalEnv.URLs.Project,
			ownername,
			reponame,
			"src",
			"master",
			entry.Name,
		)
		result = append(result, data)
	}

	return result, nil
}
