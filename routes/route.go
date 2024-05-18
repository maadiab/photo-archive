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

	router.HandleFunc("/register", middleware.Authenticate(Handlers.Signup)).Methods("POST")
	// router.HandleFunc("/addimage", middleware.Authenticate(Handlers.Addimage)).Methods("POST")

	router.HandleFunc("/login", middleware.Login).Methods("POST")
	// router.HandleFunc("/signup", Handlers.Signup).Methods("POST")

	// router.HandleFunc("/static/", Handlers.ServeForbidden)

	return router
}
