package main

import (
	"net/http"
)

const USERNAME = "username"
const PASSWORD = "password"

func Auth(w http.ResponseWriter, r *http.Request) bool  {
	username, password, ok := r.BasicAuth()

	if !ok {
		w.Write([]byte(`something went wrong`))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)

	if !isValid {
		w.Write([]byte(`wrong username or password`))
		return false
	}

	return true
}