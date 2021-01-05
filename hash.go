package main

import (
	"crypto/rand"
	"crypto/sha256"
	"io"
	"log"

	"golang.org/x/crypto/pbkdf2"
)

const PW_SALT_BYTES = 32

//Hash the password with the salt included
func HashWithSalt(password, salt []byte) []byte {
	dk := pbkdf2.Key(password, salt, 50000, 256, sha256.New)
	return dk
}

//Generate a random salt
func GenerateSalt() []byte {
	salt := make([]byte, PW_SALT_BYTES)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		log.Fatal(err)
	}

	return salt
}
