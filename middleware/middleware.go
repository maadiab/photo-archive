package middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

var JwtKey = []byte("secret_key")

var UserPerms []string

func ComparePassword(hashedPassword []byte, inputPassword string) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(inputPassword))

}

// for jwt

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username    string   `json:"username"`
	Permissions []string `json:"permissions"`
	jwt.StandardClaims
}

// check user

func CheckUser(ctx context.Context, db *sqlx.DB, user Credentials) {

}

func Login(w http.ResponseWriter, r *http.Request) {
	var cred Credentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error deconing data requested for login !!!", err)
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

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		Claims := &Claims{}

		// tkn,err:=

		if !tkn.valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

}
