package Database

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	Host     = "127.0.0.1"
	User     = "postgres"
	DbName   = "aldifaa_archive_db"
	Password = "Aa@123"
)

var DB *sqlx.DB

func ConnectDB(ctx context.Context) (*sqlx.DB, error) {
	var connStr = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", Host, User, Password, DbName)
	db, err := sqlx.ConnectContext(ctx, "postgres", connStr)

	if err != nil {
		log.Println("Connection Error !!!", err)

	}

	err = db.Ping()
	if err != nil {
		// log.Println("Ping Databasego  Error !!!", err)
		panic(err)
	}

	log.Println("connected sucsessfully ...")

	DB = db

	return db, err
}

func CreateUsersTable() error {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR (255),
		mobile VARCHAR (255),
		email VARCHAR (255),
		rank VARCHAR (255),
		hashedPassword VARCHAR (255)
	)`)

	if err != nil {
		log.Println("Error Creating users Table !!!")
	} else {
		log.Println("Creating Users Table")
	}

	return err
}

func CreatePhotographerTable() error {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS photographers (
		id SERIAL PRIMARY KEY,
		name VARCHAR (255),
		username VARCHAR (255),
		hashedPassword VARCHAR (255)
	)`)

	if err != nil {
		log.Println("Error Creating photographers Table !!!")
	} else {
		log.Println("Creating photographers Table ...")
	}

	return err
}
