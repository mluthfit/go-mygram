package helpers

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPass(pass string) string {
	var saltInt, err = strconv.Atoi(GetEnv("SALT_PASSWORD"))
	if err != nil {
		panic("SALT_PASSWORD env must be integer")
	}

	var password = []byte(pass)
	var hash, _ = bcrypt.GenerateFromPassword(password, saltInt)

	return string(hash)
}

func ComparePass(h, p []byte) bool {
	var hash, password = []byte(h), []byte(p)
	var err = bcrypt.CompareHashAndPassword(hash, password)

	return err == nil
}
