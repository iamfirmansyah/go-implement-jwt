package routes

import (
	"go-jwt/controllers"

	"github.com/gorilla/mux"
)

func User(router *mux.Router) {
	router.HandleFunc("/", controllers.GetUser).Methods("GET")
	// router.HandleFunc("", controllers.GetByID).Methods("PATCH")
	router.HandleFunc("/register", controllers.SignUp).Methods("POST")
	// router.HandleFunc("", controllers.Update).Methods("PUT")
	// router.HandleFunc("", controllers.Delete).Methods("DELETE")
}
