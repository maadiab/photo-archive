package routes

import (
	"github.com/gorilla/mux"
	Handlers "github.com/maadiab/aldifaapi/handlers"
	"github.com/maadiab/aldifaapi/middleware"
)

func Router() *mux.Router {

	// signupPermissions := []string{"read", "write"}
	//GetPhotoPermissions := []string{"read"}
	//DeletePermissions := []string{"read", "write", "delete"}
	//addPhotoPermissions := []string{"read", "write"}

	router := mux.NewRouter()

	router.HandleFunc("/", Handlers.ServeHome)
	router.HandleFunc("/login", middleware.Login).Methods("POST")
	/*router.Handle("/adduser", middleware.Authenticate(middleware.AuthorizationMiddleware(signupPermissions)(http.HandlerFunc(Handlers.Signup)))).Methods("POST")
	router.Handle("/addphoto", middleware.Authenticate(middleware.AuthorizationMiddleware(addPhotoPermissions)(http.HandlerFunc(Handlers.Addimage)))).Methods("POST")
	router.Handle("/photos", middleware.Authenticate(middleware.AuthorizationMiddleware(GetPhotoPermissions)(http.HandlerFunc(Handlers.GetPhotos)))).Methods("GET")
	router.Handle("/photos/{id}", middleware.Authenticate(middleware.AuthorizationMiddleware(GetPhotoPermissions)(http.HandlerFunc(Handlers.GetPhoto)))).Methods("GET")
	router.Handle("/photographers", middleware.Authenticate(middleware.AuthorizationMiddleware(GetPhotoPermissions)(http.HandlerFunc(Handlers.GetPhotographers)))).Methods("GET")
	router.Handle("/photographers/{id}", middleware.Authenticate(middleware.AuthorizationMiddleware(GetPhotoPermissions)(http.HandlerFunc(Handlers.GetPhotographer)))).Methods("GET")
	router.Handle("/users", middleware.Authenticate(middleware.AuthorizationMiddleware(GetPhotoPermissions)(http.HandlerFunc(Handlers.GetUsers)))).Methods("GET")
	router.Handle("/users/{id}", middleware.Authenticate(middleware.AuthorizationMiddleware(GetPhotoPermissions)(http.HandlerFunc(Handlers.GetUser)))).Methods("GET")
	router.Handle("/deleteuser/{id}", middleware.Authenticate(middleware.AuthorizationMiddleware(DeletePermissions)(http.HandlerFunc(Handlers.DeleteUser)))).Methods("POST")
	router.Handle("/deletephoto/{id}", middleware.Authenticate(middleware.AuthorizationMiddleware(DeletePermissions)(http.HandlerFunc(Handlers.DeletePhoto)))).Methods("POST")
	router.Handle("/deletephotographer/{id}", middleware.Authenticate(middleware.AuthorizationMiddleware(DeletePermissions)(http.HandlerFunc(Handlers.DeletePhotographer)))).Methods("POST")
	*/
	// router.HandleFunc("/adduser", middleware.Chain(Handlers.Signup, middleware.Authenticate(), middleware.AuthorizationMiddleware(signupPermissions))).Methods("POST")

	// router.HandleFunc("/adduser",middleware.Authenticate())

	router.HandleFunc("/adduser", middleware.Authenticate(Handlers.Signup, "admin"))
	// router.HandleFunc("/addphoto", middleware.Authenticate(Handlers.Addimage, "admin"))
	// router.HandleFunc("/photos", middleware.Authenticate(Handlers.GetPhotos, "admin"))
	// router.HandleFunc("/photos/{id}", middleware.Authenticate(Handlers.GetPhoto, "admin"))
	router.HandleFunc("/photographers", middleware.Authenticate(Handlers.GetPhotographers, "admin"))
	// router.HandleFunc("/photographers/{id}", middleware.Authenticate(Handlers.GetPhotographer, "admin"))
	router.HandleFunc("/users", middleware.Authenticate(Handlers.GetUsers, "admin"))
	// router.HandleFunc("/users/{id}", middleware.Authenticate(Handlers.GetUser, "admin"))
	router.HandleFunc("/deleteuser/{id}", middleware.Authenticate(Handlers.DeleteUser, "admin"))
	// router.HandleFunc("/deletephoto/{id}", middleware.Authenticate(Handlers.DeletePhoto, "admin"))
	// router.HandleFunc("/deletephotographer/{id}", middleware.Authenticate(Handlers.DeletePhotographer, "admin"))

	// router.Handle(middleware.Authenticate(),Handlers.AddPhotographer)

	return router
}
