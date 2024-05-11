package Handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	Database "github.com/maadiab/aldifaArchive/database"
	"github.com/maadiab/aldifaArchive/helpers"
	"golang.org/x/crypto/bcrypt"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This Is Home Page ..."))
}

func ServeLogin(w http.ResponseWriter, r *http.Request) {
	ServePages(w, "login.page.html")
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {

	// Check if the request method is POST
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form:", err)
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	name := r.Form.Get("username")
	email := r.Form.Get("email")
	mobile := r.Form.Get("mobile")
	password := r.Form.Get("password")

	err = helpers.AddUser(name, email, mobile, password)

	log.Println(password)

	if err != nil {
		http.Error(w, "Error inserting user into database", http.StatusInternalServerError)
	}

	// log.Println(r.Body)
	// Insert user data into the database
	// Respond with success message or redirect to another page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ServePages(w http.ResponseWriter, tmpl string) {
	parsedTemplates, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplates.Execute(w, nil)
	if err != nil {
		log.Println("Error Parsing template", err)
	}
}

func ServeSignup(w http.ResponseWriter, r *http.Request) {
	ServePages(w, "signup.page.html")

}

func ServeDashboard(w http.ResponseWriter, r *http.Request) {
	ServePages(w, "dashboard.page.html")
}

func ServePictures(w http.ResponseWriter, r *http.Request) {
	ServePages(w, "pictures.page.html")
}

func ServeForbidden(w http.ResponseWriter, r *http.Request) {
	ServePages(w, "forbidden.page.html")
}

func SigninHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form:", err)
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	var hashedPassword1 string
	query0 := "SELECT hashedpassword FROM users where name =$1"
	err = Database.DB.Get(&hashedPassword1, query0, username)

	log.Println(username, password, hashedPassword1)

	if err != nil {
		log.Println("Please Check Username and Password !!!", err)
		// log.Println(hashedPassword)
		return
	}

	// err = middleware.ComparePassword([]byte(hashedPassword1), password)
	// if err != nil {
	// 	log.Println("Error Hashing password !!!", err)
	// }

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword1), []byte(password))
	if err == nil {
		// Passwords match
		fmt.Println("Password is correct!")
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

	} else if err == bcrypt.ErrMismatchedHashAndPassword {
		// Passwords don't match
		fmt.Println("Password is incorrect!")
	} else {
		// Other error occurred
		fmt.Println("Error:", err)
	}

	// query := "SELECT * FROM users where name =$1 and hashedpassword =$2"

	// err = Database.DB.Get(query, username, hashedPassword1)

	// if err != nil {
	// 	log.Println("Please Check Username and Password !!!", err)
	// 	// log.Println(userCred)
	// 	return
	// }

}
