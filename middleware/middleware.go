package middleware

import (
	"fmt"
	"go-jwt/config"
	"go-jwt/helper"
	"go-jwt/models/web"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// bearer := "Bearer " + r.Header["Authorization"]

		if r.Header["Authorization"] == nil {
			response := web.WebResponse{
				Data: "No Token Found",
			}
			helper.Response(w, response, 400)
			return
		}

		var mySigningKey = []byte(config.AppConfig.JWT_KEY)

		token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			response := web.WebResponse{
				Data: "Your Token has been expired",
			}
			helper.Response(w, response, 400)
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

		response := web.WebResponse{
			Data: "Not Authorized",
		}

		helper.Response(w, response, 400)
		return
	}
}
