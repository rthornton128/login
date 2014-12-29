package main

import "testing"

func TestValidPassword(t *testing.T) {
	pwd := "password"
	salt := newSalt()
	hpwd := hashPassword(pwd, salt)

	if ok := validPassword(pwd, hpwd, salt); !ok {
		t.FailNow()
	}
}
