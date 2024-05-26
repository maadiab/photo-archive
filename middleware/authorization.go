package middleware

import (
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
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

// requiredPermissions := []string{"read", "write"}?

// func AuthorizationMiddleware(requiredPermissions []string) Middleware {
// 	return func(f http.HandlerFunc) http.HandlerFunc {
// 		return func(w http.ResponseWriter, r *http.Request) {
// 			//v := r.Context().Value("claims")
// 			//fmt.Printf("v is %#v", v)
// 			//if v == nil {
// 			//log.Println("No permissions found !!!", ok)
// 			//http.Error(w, "Permission not found !!!", http.StatusInternalServerError)
// 			//return
// 			//}

// 			//log.Println(claims.Permissions)
// 			/*
// 				if !hasPermissions(claims.Permissions, requiredPermissions) {
// 					log.Println("Insufficient permission !!!")
// 					http.Error(w, "Insufficient permission !!!", http.StatusForbidden)
// 					return
// 				}*/
// 			// Call the next function
// 			f(w, r)
// 		}
// 	}
// }

// func AuthorizationMiddleware(requiredPermissions string, next http.Handler) http.HandlerFunc {

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// next.ServeHTTP(w, r)

// func(w http.ResponseWriter, r *http.Request) {
//v := r.Context().Value("claims")
//fmt.Printf("v is %#v", v)
//if v == nil {
//log.Println("No permissions found !!!", ok)
//http.Error(w, "Permission not found !!!", http.StatusInternalServerError)
//return
//}

//log.Println(claims.Permissions)
/*
	if !hasPermissions(claims.Permissions, requiredPermissions) {
		log.Println("Insufficient permission !!!")
		http.Error(w, "Insufficient permission !!!", http.StatusForbidden)
		return
	}*/
// Call the next function
// next.ServeHTTP(w, r)
// }
// 	})
// }

func AuthorizationMiddleware(requiredPermissions string) Middleware {

	return func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			cookie, err := r.Cookie("token")
			if err != nil {
				log.Println("Error: ", err)
				return
			}

			tokenStr := cookie.Value
			claims := &Claims{}

			jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
				return JwtKey, nil
			})

			// check user permissions
			if requiredPermissions != claims.Permission {
				log.Println("Error: user does not have permissions !!!")
				return
			}

			log.Println("User permissions is: ", claims.Permission)

			// if err != nil {
			// 	log.Println("Error: ", err)
			// 	return
			// }

			// fmt.Printf("v is %#v", tokenStr)

			// if v == nil {
			// 	log.Println("No permissions found !!!")
			// 	http.Error(w, "Permission not found !!!", http.StatusInternalServerError)
			// 	return
			// }

			//log.Println(claims.Permissions)
			/*
				if !hasPermissions(claims.Permissions, requiredPermissions) {
					log.Println("Insufficient permission !!!")
					http.Error(w, "Insufficient permission !!!", http.StatusForbidden)
					return
				}*/
			// Call the next function
			h.ServeHTTP(w, r)
		}
	}

}
