package main

const (
	// ErrIDNotFound means that there was an error with the ID
	ErrIDNotFound = Err("article not found")

	// ErrUsernameUnavailable means that the username is unavailable
	ErrUsernameUnavailable = Err("username isn't available")

	// ErrPasswordNotEmpty means that the password can't be empty
	ErrPasswordNotEmpty = Err("password can't be empty")
)

// Err are errors that can happen when interacting with FSPS
type Err string

// The Error func returns the FSPS Error
func (e Err) Error() string {
	return string(e)
}
