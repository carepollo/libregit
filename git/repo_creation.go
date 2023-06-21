package git

import (
	"fmt"
	"os"
	"time"

	"github.com/carepollo/librecode/utils"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

// create a directory in given path and initialize a server-side repository
// in the git root location
func CreateRepo(path string) (*git.Repository, error) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return nil, err
	}

	return git.PlainInit(path, true)
}

// adds the default README for repo on the root of the repo's project directory
func AddReadme(path string) error {
	url := fmt.Sprintf("%v/%v", utils.GlobalEnv.URLs.Project, path)
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		return err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	readmeLocation := fmt.Sprintf("%v/README.md", utils.GlobalEnv.ProjectRoot)
	_, err = worktree.Add(readmeLocation)
	if err != nil {
		return err
	}

	_, err = worktree.Commit("User registration", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "librecode",
			Email: "system@librecode.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		return err
	}

	return repo.Push(&git.PushOptions{
		RemoteName: "origin",
	})
}
