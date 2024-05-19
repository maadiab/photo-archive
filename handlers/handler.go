package Handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/maadiab/aldifaapi/core"
	Database "github.com/maadiab/aldifaapi/database"
	"github.com/maadiab/aldifaapi/helpers"
	// "text/template"
)

// func hasPermissions(userPermissions []string, requiredPermissions []string) bool {
// 	for _, perm := range requiredPermissions {
// 		found := false
// 		for _, userPerm := range userPermissions {
// 			if perm == userPerm {
// 				found = true
// 				break
// 			}
// 		}
// 		if !found {
// 			return false
// 		}
// 	}
// 	return true
// }

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This Is Home Page ..."))
}

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is signup page ..."))
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// requiredPermissions := []string{"read", "write"}
	// claims, ok := r.Context().Value("claims").(*middleware.Claims)
	// if !ok {
	// 	log.Println("No permissions found !!!", ok)
	// 	http.Error(w, "Permission not found !!!", http.StatusInternalServerError)
	// 	return
	// }

	// log.Println(claims.Permissions)
	// if !hasPermissions(claims.Permissions, requiredPermissions) {
	// 	log.Println("nsuffecient permission !!!")
	// 	http.Error(w, "Insuffecient permission !!!", http.StatusForbidden)
	// 	return
	// }

	var user core.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	helpers.AddUser(Database.DB, user)

}

func Addimage(w http.ResponseWriter, r *http.Request) {
	var photo core.Photo

	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		log.Println("Error decoding photo !!!", err)
		return
	}

	helpers.Addimage(Database.DB, photo)
}

func GetPhoto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 32, 32)
	if err != nil {
		log.Println("Error: no user found !!!")
		return
	}

	user, nil := helpers.GetPhoto(Database.DB, int(userID))
	jsonData, err := json.Marshal(user)

	if err != nil {
		log.Println("Error: no user found !!!", err)
	}

	w.Write(jsonData)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["id"], 32, 32)
	if err != nil {
		log.Println("Error: no user found !!!", err)
		return
	}
	user, nil := helpers.GetUser(Database.DB, int(userId))
	jsonData, err := json.Marshal(user)

	if err != nil {
		log.Println("Error: no user found !!!", err)
	}

	w.Write(jsonData)

}

func GetPhotographer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 32, 32)
	if err != nil {
		log.Println("Error: no photographer found !!!", err)
	}

	photographer, nil := helpers.GetPhotographer(Database.DB, int(userID))
	if err != nil {
		log.Println("Error: no photographer with given id found !!! ", err)

	}

	jsonData, err := json.Marshal(photographer)

	if err != nil {
		log.Println("Error marshalling photographer !!!", err)
	}
	w.Write(jsonData)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 32, 32)

	if err != nil {
		log.Println("Error: cannot delete this user, not found !!!", err)

	}

	err = helpers.DeleteUser(Database.DB, int(userID))

	if err != nil {
		log.Println("Error: ", err)
	}
	w.Write([]byte("User Deleted successfully ..."))

	log.Println("User deleted successfully ...")
}

// func ServePages(w http.ResponseWriter, tmpl string) {
// 	parsedTemplates, _ := template.ParseFiles("./templates/" + tmpl)
// 	err := parsedTemplates.Execute(w, nil)
// 	if err != nil {
// 		log.Println("Error Parsing template", err)
// 	}
// }
