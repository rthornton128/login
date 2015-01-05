package crypt_test

import (
	"testing"

	"github.com/rthornton128/login/crypt"
)

func TestValidPassword(t *testing.T) {
	raw_password := "password"
	hashed_password, salt := crypt.Hash(raw_password)

	if ok := crypt.Validate(raw_password, hashed_password, salt); !ok {
		t.FailNow()
	}
}
