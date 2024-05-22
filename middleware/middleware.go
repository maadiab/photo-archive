package middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
	"github.com/maadiab/aldifaapi/core"
	Database "github.com/maadiab/aldifaapi/database"
	"golang.org/x/crypto/bcrypt"
)

var JwtKey = []byte("secret_key")

var UserPerms []string

func ComparePassword(hashedPassword []byte, inputPassword string) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(inputPassword))

}

// check user
func CheckUser(ctx context.Context, db *sqlx.DB, user Credentials) error {
	var userCred core.User
	var hashedPassword string

	// _, err := Database.DB.Exec("SELECT hashedPassword FORM users where username = $1", user.Username)
	err := db.Get(&hashedPassword, "SELECT hashedPassword FROM users WHERE username =$1", user.Username)

	if err != nil {
		log.Println("Please check username and password !!!", err)
		return err
	}

	err = ComparePassword([]byte(hashedPassword), user.Password)
	if err != nil {
		log.Println("Please check your password !!!", err)
		return err
	}

	err = db.Get(&userCred, "SELECT * FROM users WHERE username = $1 ", user.Username)
	if err != nil {
		log.Println("Error: ", err)
		return err
	}

	var userPermissions []string
	err = db.Select(&userPermissions, "SELECT permission_type FROM permissions WHERE user_type =$1", userCred.Permissions)
	if err != nil {
		log.Println("Error: no permissions found !!!", err)
		return err
	}

	UserPerms = userPermissions

	log.Println("user verified successfully ...")

	log.Println("user credential: ", userCred.Username)
	log.Println(UserPerms)
	log.Println(userPermissions)

	return err
}

// Give user Jwt token

// jwt main structs
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username    string   `json:"username"`
	Permissions []string `json:"permissions"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	var cred Credentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error deconing data requested for login !!!", err)
		return
	}

	// check user
	err = CheckUser(r.Context(), Database.DB, cred)
	if err != nil {
		log.Println("Error: authentication error !!! ")
		return
	}

	// Setting claims
	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &Claims{
		Username:    cred.Username,
		Permissions: UserPerms,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Authentication Error !!!", err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

// Refresh token

func RefreshToken(w http.ResponseWriter, r *http.Request) {
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
	expirationTime := time.Now().Add(time.Minute * 15)

	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		log.Println("Error: ", err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "refreshed_token",
		Value:   tokenString,
		Expires: expirationTime,
	})

}

// finish refresh

// Authentications (real middleware)
// func Authenticate(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		cookie, err := r.Cookie("token")
// 		if err != nil {
// 			if err == http.ErrNoCookie {
// 				w.WriteHeader(http.StatusUnauthorized)
// 				return
// 			}
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}

// 		tokenStr := cookie.Value
// 		Claims := &Claims{}

// 		tkn, err := jwt.ParseWithClaims(tokenStr, Claims, func(t *jwt.Token) (interface{}, error) {
// 			return JwtKey, nil
// 		})

// 		if err != nil {
// 			if err == jwt.ErrSignatureInvalid {
// 				w.WriteHeader(http.StatusUnauthorized)
// 				return
// 			}
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}

// 		if !tkn.Valid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}

// 		log.Println("Hello, ", Claims.Username)

// 		// ctx := context.WithValue(r.Context(), "Claims", Claims)
// 		ctx := context.WithValue(r.Context(), "claims", Claims)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	}

// }

// modified version from middleware
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenStr := cookie.Value
		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		log.Println("Authenticated user:", claims.Username)

		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
