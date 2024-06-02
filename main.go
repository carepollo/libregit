package main

import "github.com/carepollo/noxt/router"

func main() {
	tree := router.NewTree()

	tests := []string{
		"/this/is/a/test",
		"/this/is/a/possible/test",
	}

	for _, v := range tests {
		tree.Insert(v)
	}
}
