package middleware

//To check ...

// func checkRole(requiredRole string, next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         // Check if user has the required role
//         // If authorized, call the next handler
//         // Otherwise, return an error response
//         next.ServeHTTP(w, r)
//     })
// }

// // Usage:
// router := mux.NewRouter()
// router.Use(validateToken)
// router.HandleFunc("/admin", checkRole("admin", adminHandler))
