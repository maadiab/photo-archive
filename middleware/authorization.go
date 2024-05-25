package middleware

import (
	"net/http"
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

func AuthorizationMiddleware(requiredPermissions []string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
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
			f(w, r)
		}
	}
}
