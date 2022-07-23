package main

import (
	"fmt"
	"go-jwt/app"
	"go-jwt/config"
	"go-jwt/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func main() {
	// * Load Configuration
	config.LoadAppConfig()

	// * Initialize Database
	app.ConnectMysql(config.AppConfig.MYSQL_CONNECTION)

	// * Initialize the router
	CreateRouter()

	// * Register Routes
	InitializeRoute()

	// * Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", config.AppConfig.PORT))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.AppConfig.PORT), router))
}

func CreateRouter() {
	router = mux.NewRouter()
}

func InitializeRoute() {
	s := router.PathPrefix("/api").Subrouter()
	routes.User(s.PathPrefix("/user").Subrouter())
}
