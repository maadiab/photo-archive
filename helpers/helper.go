package helpers

import (
	"log"

	Database "github.com/maadiab/aldifaArchive/database"
	"golang.org/x/crypto/bcrypt"
)

func HashingPasswords(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func AddUser(username string, email string, mobile string, password string) error {

	hashedPassword, err := HashingPasswords(password)
	if err != nil {
		log.Println("Error hashing password !!!")
	}
	// log.Println(username, email, mobile, hashedPassword)

	_, err = Database.DB.Exec("INSERT INTO users (name, email, mobile, hashedpassword) VALUES ($1, $2, $3, $4)", username, email, mobile, hashedPassword)

	if err != nil {
		log.Println("Error inserting user into database:", err)

		return err
	}

	log.Println("User added successfully ...")
	return err
}

// Serve HTML Templates
