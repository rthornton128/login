package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha512"
	"log"
)

const size = sha512.Size

func newSalt() []byte {
	salt := make([]byte, size)
	if _, err := rand.Read(salt); err != nil {
		log.Println("new salt:", err)
	}
	return salt
}

func hashPassword(pwd string, salt []byte) []byte {
	h := sha512.Sum512([]byte(pwd + string(salt)))
	return h[:]
}

func validPassword(pwd string, hpwd, salt []byte) bool {
	h := sha512.Sum512([]byte(pwd + string(salt)))
	return bytes.Equal(h[:], hpwd)
}
