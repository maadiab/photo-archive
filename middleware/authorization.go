package middleware

import (
	"log"
	"net/http"
)

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

// requiredPermissions := []string{"read", "write"}?

func AuthorizationMiddleware(requiredPermissions []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value("claims").(*Claims)
			if !ok {
				log.Println("No permissions found !!!", ok)
				http.Error(w, "Permission not found !!!", http.StatusInternalServerError)
				return
			}

			log.Println(claims.Permissions)
			if !hasPermissions(claims.Permissions, requiredPermissions) {
				log.Println("Insufficient permission !!!")
				http.Error(w, "Insufficient permission !!!", http.StatusForbidden)
				return
			}
			// Call the next function
			next.ServeHTTP(w, r)

		})
	}

}
