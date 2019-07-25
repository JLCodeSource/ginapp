package main

import (
	"strings"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

// For this demo, we're storing the user list in memory
// We also have some users predefined.
// In a real application, this list will most likely be fetched
// from a database. Moreover, in production settings, you should
// store passwords securely by salting and hashing them instead
// of using them as we're doing in this demo
// TODO - convert list to test & create a persistent store for userList
var userList = []user{
	{Username: "user1", Password: "pass1"},
	{Username: "user2", Password: "pass2"},
	{Username: "user3", Password: "pass3"},
}

func registerNewUser(username, password string) (*user, error) {

	if strings.TrimSpace(password) == "" {
		return nil, ErrPasswordNotEmpty
	} else if !isUsernameAvailable(username) {
		return nil, ErrUsernameUnavailable
	}

	u := user{Username: username, Password: password}

	userList = append(userList, u)

	return &u, nil

}

func isUsernameAvailable(username string) bool {
	for _, u := range userList {
		if u.Username == username {
			return false
		}
	}
	return true
}

func isUserValid(username, password string) bool {
	for _, u := range userList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}
