package helpers

import (
	"errors"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPass(pass string) string {
	var saltInt, err = strconv.Atoi(os.Getenv("SALT_PASSWORD"))
	if err != nil {
		panic(errors.New("SALT_PASSWORD env must be integer"))
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
