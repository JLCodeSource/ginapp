package main

import (
	"testing"
)

func TestUsernameAvailability(t *testing.T) {
	saveLists()
	newusername := "newuser"
	existingusername := "user1"

	if !isUsernameAvailable(newusername) {
		t.Errorf("expected username '%s' to be available, but it is not", newusername)
	}

	if isUsernameAvailable(existingusername) {
		t.Errorf("expected username '%s' to be unavailable, but it is not", existingusername)
	}

	registerNewUser("newuser", "newpass")
	if isUsernameAvailable(newusername) {
		t.Errorf("expected username '%s' to be unavailable, but it is not", newusername)
	}

	restoreLists()
}

func TestValidUserRegistration(t *testing.T) {
	saveLists()

	u, err := registerNewUser("newuser", "newpass")
	empty := ""

	if err != nil {
		t.Errorf("did not expect an error but got '%s'", err)
	}

	if u.Username == empty {
		t.Errorf("wanted username '%s' but got '%s'", u.Username, empty)
	}
}

func TestInvalidUserRegistration(t *testing.T) {
	saveLists()

	u, err := registerNewUser("user1", "pass1")

	if err == nil {
		t.Errorf("expected error but got none")
	}

	if u != nil {
		t.Errorf("expected nil response, but got '%s'", u)
	}

	u, err = registerNewUser("newuser", "")

	
	if err == nil {
		t.Errorf("expected error but got none")
	}

	if u != nil {
		t.Errorf("expected nil response, but got '%s'", u)
	}

	restoreLists()
}