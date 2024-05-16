package Handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/maadiab/aldifaapi/core"
	Database "github.com/maadiab/aldifaapi/database"
	"github.com/maadiab/aldifaapi/helpers"
	"github.com/maadiab/aldifaapi/middleware"
	// "text/template"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This Is Home Page ..."))
}

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is signup page ..."))
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	requiredPermissions := []string{"read", "write"}
	claims, ok := r.Context().Value("claims").(*middleware.Claims)
	if !ok {
		log.Println("No permissions found !!!", ok)
		http.Error(w, "Permission not found !!!", http.StatusInternalServerError)
		return
	}

	log.Println(claims.Permissions)
	if !hasPermissions(claims.Permissions, requiredPermissions) {
		log.Println("nsuffecient permission !!!")
		http.Error(w, "Insuffecient permission !!!", http.StatusForbidden)
		return
	}

	var user core.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	helpers.AddUser(Database.DB, user)

}

func ServePages(w http.ResponseWriter, tmpl string) {
	parsedTemplates, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplates.Execute(w, nil)
	if err != nil {
		log.Println("Error Parsing template", err)
	}
}

func hasPermissions(userPermissions []string, requiredPermissions []string) bool {
	for _, perm := range requiredPermissions {
		found := false
		for _, userPerm := range userPermissions {
			if perm == userPerm {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
