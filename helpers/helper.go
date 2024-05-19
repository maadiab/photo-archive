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

	_, err = Database.DB.Exec("INSERT INTO users (name, username,email, mobile, hashedPassword,permission_type) VALUES ($1, $2, $3, $4,$5,$6)", user.Name, user.Username, user.Email, user.Mobile, hashedPassword, user.Permissions)

	if err != nil {
		log.Println("Error inserting user into database:", err)

		return err
	}

	log.Println("User added successfully ...")
	return err
}

func Addimage(db *sqlx.DB, photo core.Photo) error {
	_, err := Database.DB.Exec("INSERT INTO photos (name, photographer, tags) VALUES ($1,$2,$3)", photo.Name, photo.Photographer, photo.Tags)
	if err != nil {
		log.Println("Error inserting photo into database: ", err)
		return err
	}
	log.Println("photo added successfully ...")
	return err
}

func GetPhoto(db *sqlx.DB, userID int) (core.Photo, error) {

	var photo core.Photo

	err := db.Get(&photo, "SELECT * FROM photos WHERE id =$1", userID)
	if err != nil {
		log.Println("Error: getting photo !!!", err)

	}
	return photo, err
}

func GetUser(db *sqlx.DB, userID int) (core.User, error) {

	var user core.User

	err := db.Get(&user, "SELECT * FROM users WHERE id = $", userID)
	if err != nil {
		log.Println("Error: no user found !!!", err)

	}
	return user, err

}

func GetPhotographer(db *sqlx.DB, userID int) (core.Photographer, error) {
	var photographer core.Photographer

	err := db.Get(&photographer, "SELECT FROM photographer WHERE id = $1", userID)

	if err != nil {
		log.Println("Error: no photographer found !!!", err)
	}

	return photographer, err

}

func DeleteUser(db *sqlx.DB, userID int) error {

	_, err := db.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		log.Println("Error: Cannot delete this user !!!", err)
		return err
	}
	return err
}

func DeletePhoto(db *sqlx.DB, userID int) error {
	_, err := db.Exec("DELETE * FROM photos WHERE id = $1", userID)
	if err != nil {
		log.Println("Error: Cannot delete this user !!!", err)
		return err
	}
	return err
}

func DeletePhotographer(db *sqlx.DB, userID int) error {

	_, err := db.Exec("DELETE FROM photographers WHERE id =$1", userID)
	if err != nil {
		log.Println("Error: Cannot delete this photographer !!!", err)
		return err
	}

	return err
}

func GetUsers(db *sqlx.DB) ([]core.User, error) {
	var users []core.User

	err := db.Select(&users, "SELECT * FROM users")
	if err != nil {
		log.Println("Error: ", err)
		return nil, err
	}
	return users, err
}

func Getphotos(db *sqlx.DB) ([]core.Photo, error) {

	var photos []core.Photo

	err := db.Select(&photos, "SELECT * FROM photos")
	if err != nil {
		log.Println("Error: ", err)
		return nil, err
	}
	return photos, err
}

func GetPhotographers(db *sqlx.DB) ([]core.Photographer, error) {

	var photographers []core.Photographer

	err := db.Select(&photographers, "SELECT * FROM photographers")
	if err != nil {

		log.Println("Error: ", err)
		return nil, err
	}
	return photographers, err

}
