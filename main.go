package main

import (
	"fmt"
	"net/http"
)

func main() {
	protectedHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, r.URL.Path+" este es el path xd")
	}

	authHandler := func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if !ok || username != "admin" && password != "test" {
			w.Header().Set("WWW-Authenticate", `Basic Realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "not right credentials")
			return
		}

		protectedHandler(w, r)
	}

	http.HandleFunc("/", authHandler)

	http.ListenAndServe(":5000", nil)
}
