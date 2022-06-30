package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(plainPassword string) (hashedPassword string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 10)
	return string(bytes), err
}

func CompareHashPassword(plainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
