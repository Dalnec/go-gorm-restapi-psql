package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Dalnec/go-gorm-restapi-psql/helpers"
	"github.com/golang-jwt/jwt"
)

//check whether user is authorized or not
func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		if r.Header["Authorization"] == nil {
			var err helpers.Error
			err = helpers.SetError(err, "No Token Found")
			json.NewEncoder(w).Encode(err)
			return
		}
		
		secretkey := os.Getenv("SECRET")
		var mySigningKey = []byte(secretkey)

		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]
		token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing token.")
			}
			return mySigningKey, nil
		})

		if err != nil {
			var err helpers.Error
			err = helpers.SetError(err, "Your Token has been expired.")
			json.NewEncoder(w).Encode(err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {
				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return

			} else if claims["role"] == "user" {
				r.Header.Set("Role", "user")
				handler.ServeHTTP(w, r)
				return

			}
		}
		var reserr helpers.Error
		reserr = helpers.SetError(reserr, "Not Authorized.")
		json.NewEncoder(w).Encode(err)
	}
}