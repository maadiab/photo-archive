package routes

import (
	"github.com/gorilla/mux"
	Handlers "github.com/maadiab/aldifaapi/handlers"
	"github.com/maadiab/aldifaapi/middleware"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	// router.HandleFunc("/", Handlers.ServeLogin)
	// router.HandleFunc("/login", Handlers.ServeLogin)
	router.HandleFunc("/register", Handlers.ServeSignup).Methods("GET")
	router.HandleFunc("/signup", Handlers.SignupHandler).Methods("POST")
	router.HandleFunc("/signin", Handlers.SigninHandler).Methods("POST")

	router.HandleFunc("/dashboard", Handlers.ServeDashboard)
	router.HandleFunc("/pictures", Handlers.ServeDashboard)

	router.HandleFunc("/login", middleware.Login)
	// router.HandleFunc("/static/", Handlers.ServeForbidden)

	return router
}
