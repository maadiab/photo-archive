package helpers

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/maadiab/aldifaapi/core"
	Database "github.com/maadiab/aldifaapi/database"
	"golang.org/x/crypto/bcrypt"
)

func HashingPasswords(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// func AddUser(username string, email string, mobile string, password string) error {

// 	hashedPassword, err := HashingPasswords(password)
// 	if err != nil {
// 		log.Println("Error hashing password !!!")
// 	}
// 	// log.Println(username, email, mobile, hashedPassword)

// 	_, err = Database.DB.Exec("INSERT INTO users (name, email, mobile, hashedpassword) VALUES ($1, $2, $3, $4)", username, email, mobile, hashedPassword)

// 	if err != nil {
// 		log.Println("Error inserting user into database:", err)

// 		return err
// 	}

// 	log.Println("User added successfully ...")
// 	return err
// }

func AddUser(db *sqlx.DB, user core.User) error {

	hashedPassword, err := HashingPasswords(user.Password)
	if err != nil {
		log.Println("Error hashing password !!!")
	}
	// log.Println(username, email, mobile, hashedPassword)

	_, err = Database.DB.Exec("INSERT INTO users (name, username,email, mobile, hashedpassword) VALUES ($1, $2, $3, $4,$5)", user.Name, user.Username, user.Email, user.Mobile, hashedPassword)

	if err != nil {
		log.Println("Error inserting user into database:", err)

		return err
	}

	log.Println("User added successfully ...")
	return err
}
