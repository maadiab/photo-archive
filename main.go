package main

import (
	"context"
	"log"
	"net/http"

	Database "github.com/maadiab/aldifaArchive/database"
	Handlers "github.com/maadiab/aldifaArchive/handlers"
	"github.com/maadiab/aldifaArchive/routes"
)

// func init() {
// 	fs := http.FileServer(http.Dir("static"))
// 	http.Handle("/static/", http.StripPrefix("static", fs))

// 	storage := http.FileServer(http.Dir("storage"))
// 	http.Handle("/storage/", http.StripPrefix("storage", storage))
// }

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

	r := routes.Router()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/login", Handlers.ServeLogin)

	log.Println("server is running at port: 8080 ...")
	http.ListenAndServe(":8080", r)

	defer Database.DB.Close()

}
