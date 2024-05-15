package routes

import (
	"github.com/gorilla/mux"
	Handlers "github.com/maadiab/aldifaapi/handlers"
	"github.com/maadiab/aldifaapi/middleware"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", Handlers.ServeHome)
	// router.HandleFunc("/login", Handlers.ServeLogin)

	router.HandleFunc("/signup", middleware.Authenticate(Handlers.Signup))

	// router.HandleFunc("/static/", Handlers.ServeForbidden)

	return router
}
