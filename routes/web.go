package routes

import (
	"go-jwt/controllers"

	"github.com/gorilla/mux"
)

func User(router *mux.Router) {
	router.HandleFunc("/register", controllers.SignUp).Methods("POST")
	router.HandleFunc("/login", controllers.SignIn).Methods("POST")
}
