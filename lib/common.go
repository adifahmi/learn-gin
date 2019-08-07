package lib

import (
	// "fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

// JSON type
type JSON = map[string]interface{}

// Hash generate bcrypt of a string
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

// CheckHash will check if hash is correct
func CheckHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateToken will generate token from json
func GenerateToken(data JSON, days int) (string, error) {

	//  token is valid for X days
	date := time.Now().Add(time.Hour * 24 * time.Duration(days))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": data,
		"exp":  date.Unix(),
	})

	key := []byte(os.Getenv("SECRET_KEY"))
	tokenString, err := token.SignedString(key)
	return tokenString, err
}

// ParseToken accept token string and decode it
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(
		tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		},
	)
	return claims, err
}
