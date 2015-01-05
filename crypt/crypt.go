/* Crypto implements a simple password encryption method. bcrypt would be
 * better in a production environment but crypto/rand will more suffice for
 * this example (and probably even in many live programs, too) */
package crypt

import (
	"bytes"
	"crypto/rand"
	"crypto/sha512"
	"log"
)

/* Hash generates a sha512 password hash and a sha512.Size salt */
func Hash(raw_password string) (password, salt []byte) {
	/* generate salt (uses /dev/urandom on Linux). Only a half-hearted attempt
	 * is made to track an error. Note that Read can return less than 16 bytes,
	 * which would make for a poor salt, but is not checked */
	salt = make([]byte, sha512.Size)
	if _, err := rand.Read(salt); err != nil {
		log.Println("new salt:", err)
	}

	/* hash the password using salt and return both values */
	hashed_password := sha512.Sum512([]byte(raw_password + string(salt)))
	return hashed_password[:], salt
}

/* Validate returns true if the supplied plain text password hashes to
 * the supplied hashed password and salt combination; otherwise, false */
func Validate(raw_password string, hashed_password, salt []byte) bool {
	hashed_result := sha512.Sum512([]byte(raw_password + string(salt)))
	return bytes.Equal(hashed_result[:], hashed_password)
}
