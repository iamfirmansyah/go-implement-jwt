package routes

import (
	"go-jwt/controllers"
	"go-jwt/middleware"

	"github.com/gorilla/mux"
)

func User(router *mux.Router) {
	router.HandleFunc("/register", controllers.SignUp).Methods("POST")
	router.HandleFunc("/login", controllers.SignIn).Methods("POST")
	router.HandleFunc("/admin-page", middleware.IsAuthorized(controllers.AdminIndex)).Methods("GET")
	router.HandleFunc("/user-page", middleware.IsAuthorized(controllers.UserIndex)).Methods("GET")

}
