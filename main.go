package main

import (
	"context"
	"log"
	"net/http"

	Database "github.com/maadiab/aldifaapi/database"
	"github.com/maadiab/aldifaapi/routes"
)

func main() {
	ctx := context.Background()
	_, err := Database.ConnectDB(ctx)
	if err != nil {
		log.Println("Error Connecting Database !!!", err)
	}
	err = Database.CreateUsersTable()
	if err != nil {
		log.Println("main: Error Creating users Table !!!", err)
	}
	err = Database.CreatePhotographerTable()
	if err != nil {
		log.Println("main: Error Creating photographers Table !!!", err)
	}
	err = Database.CreatePhotosTable()
	if err != nil {
		log.Println("main: Error Creating phogtos Table !!!", err)
	}
	err = Database.CreatePermissionsTable()
	if err != nil {
		log.Println("main: Error Creating permissions Table !!!", err)
	}
	r := routes.Router()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("server is running at port: 8080 ...")
	http.ListenAndServe(":8080", r)
	defer Database.DB.Close()
}
