package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword converts current string password to hashed byte password
func HashPassword(password string) ([]byte, error) {
	bytePass := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)

	if err != nil {
		log.Println("Error in generatig hashed password: " + err.Error())
		return nil, err
	}

	return hashedPassword, nil
}

// ComparePassword checks if input matched the password in DB
func ComparePassword(hashedPassword []byte, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
