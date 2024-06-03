package main

import (
	"fmt"

	"github.com/carepollo/noxt/router"
)

func main() {
	tree := router.NewTree()
	values := []string{
		"/this/is/a/test",
		"/this/is/a/possible/test",
		"/this/path/is/not/here",
		"/this/is/a/potential/path",
	}

	for _, v := range values {
		tree.Insert(v)
	}

	tests := []string{
		"",
	}

	for _, v := range tests {
		fmt.Printf("%s was found: %v", v, tree.Search(v) != nil)
	}
}
