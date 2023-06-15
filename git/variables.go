package git

import "github.com/go-git/go-git/v5/plumbing/transport/server"

var (
	Gitserver = server.DefaultServer // where all the repos will be stored
	GitPath   = ""                   // git server instance of library
)
