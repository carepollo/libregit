package controllers

import "github.com/go-git/go-git/v5/plumbing/transport/server"

var (
	gitserver = server.DefaultServer // git server instance of library
	gitpath   = "/tmp"               // where all the repos will be stored
)
