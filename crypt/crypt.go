// Package crypt implements a simple password encryption method. Using
// SHA512 encryption is potentially dangerous. A more secure hashing algorithm
// like bcrypt (https://godoc.org/golang.org/x/crypto/bcrypt) or scrypt
// (https://godoc.org/golang.org/x/crypto/scrypt) should be used for anything
// intended to be used in a live environment. For the purpose of this
// program, the cryptography algorithms used will suffice.
package crypt

import (
	"crypto/rand"
	"crypto/sha512"
	"crypto/subtle"
	"log"
)

// Hash generates a sha512 password hash and a sha512.Size salt
func Hash(rawPassword string) (password, salt []byte) {
	// generate salt (uses /dev/urandom on Linux). Only a half-hearted attempt
	// is made to track an error. Note that Read can return less than 16 bytes,
	// which would make for a poor salt, but is not checked
	salt = make([]byte, sha512.Size)
	if _, err := rand.Read(salt); err != nil {
		log.Println("new salt:", err)
	}

	// hash the password using salt and return both values
	hashedPassword := sha512.Sum512([]byte(rawPassword + string(salt)))
	return hashedPassword[:], salt
}

// Validate returns true if the supplied plain text password hashes to
// the supplied hashed password and salt combination; otherwise, false
func Validate(rawPassword string, hashedPassword, salt []byte) bool {
	hashedResult := sha512.Sum512([]byte(rawPassword + string(salt)))
	return subtle.ConstantTimeCompare(hashedResult[:], hashedPassword) == 1
}
