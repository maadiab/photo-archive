package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	Handlers "github.com/maadiab/aldifaapi/handlers"
	"github.com/maadiab/aldifaapi/middleware"
)

func Router() *mux.Router {

	signupPermissions := []string{"read", "write"}
	GetPhotoPermissions := []string{"read", "write"}
	DeletePermissions := []string{"read", "write", "delete"}
	// GetPhotographerPermissions := []string{"read", "write"}

	router := mux.NewRouter()

	router.HandleFunc("/", Handlers.ServeHome)
	// router.HandleFunc("/login", Handlers.ServeLogin)

	// router.HandleFunc("/register", middleware.Authenticate(Handlers.Signup)).Methods("POST")
	// router.HandleFunc("/register", middleware.Authenticate(middleware.AuthorizationMiddleware(signupPermissions)(http.HandlerFunc(Handlers.Signup)))).Methods("POST")

	router.Handle("/signup", middleware.Authenticate(middleware.AuthorizationMiddleware(signupPermissions)(http.HandlerFunc(Handlers.Signup)))).Methods("POST")
	router.Handle("/photos", middleware.Authenticate(middleware.AuthorizationMiddleware(GetPhotoPermissions)(http.HandlerFunc(Handlers.GetPhoto)))).Methods("GET")
	router.Handle("/photos/{id}", middleware.Authenticate(middleware.AuthorizationMiddleware(GetPhotoPermissions)(http.HandlerFunc(Handlers.GetPhoto)))).Methods("GET")
	router.Handle("/photographers", middleware.Authenticate(middleware.AuthorizationMiddleware(GetPhotoPermissions)(http.HandlerFunc(Handlers.GetPhotographer)))).Methods("GET")
	router.Handle("/photographers/{id}", middleware.Authenticate(middleware.AuthorizationMiddleware(GetPhotoPermissions)(http.HandlerFunc(Handlers.GetPhoto)))).Methods("GET")
	router.Handle("/users", middleware.Authenticate(middleware.AuthorizationMiddleware(GetPhotoPermissions)(http.HandlerFunc(Handlers.GetUsers)))).Methods("GET")
	router.Handle("/users/{id}", middleware.Authenticate(middleware.AuthorizationMiddleware(GetPhotoPermissions)(http.HandlerFunc(Handlers.GetUser)))).Methods("GET")

	router.Handle("/deleteuser/{id}", middleware.Authenticate(middleware.AuthorizationMiddleware(DeletePermissions)(http.HandlerFunc(Handlers.DeleteUser)))).Methods("POST")
	router.Handle("/deleteuser/{id}", middleware.Authenticate(middleware.AuthorizationMiddleware(DeletePermissions)(http.HandlerFunc(Handlers.DeleteUser)))).Methods("POST")

	// router.Handle("/photographer/{id}", middleware.Authenticate(middleware.AuthorizationMiddleware(GetPhotoPermissions)(http.HandlerFunc(Handlers.)))).Methods("GET")

	// router.HandleFunc("/addimage", middleware.Authenticate(Handlers.Addimage)).Methods("POST")

	router.HandleFunc("/login", middleware.Login).Methods("POST")
	// router.HandleFunc("/signup", Handlers.Signup).Methods("POST")

	// router.HandleFunc("/static/", Handlers.ServeForbidden)

	return router
}
