package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	bytePass := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)

	if err != nil {
		log.Println("Error in generatig hashed password: " + err.Error())
		return nil, err
	}

	return hashedPassword, nil
}
