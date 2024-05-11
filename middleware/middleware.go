package middleware

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)

func ComparePassword(hashedPassword []byte, inputPassword string) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(inputPassword))

}



