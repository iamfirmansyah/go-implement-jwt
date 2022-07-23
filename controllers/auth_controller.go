package controllers

import (
	"go-jwt/app"
	"go-jwt/helper"
	"go-jwt/models/domain"
	"go-jwt/models/web"
	"go-jwt/services"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()

	var user web.UserRequest

	helper.ReadFromRequestBody(r, &user)

	err := validate.Struct(user)

	errors := helper.FormatValidationError(err)

	if errors != nil {
		response := web.WebResponse{
			Data: errors,
		}
		helper.Response(w, response, 400)
		return
	}

	var dbuser domain.User
	app.Instance.Where("email = ?", user.Email).First(&dbuser)

	//checks if email is already register or not
	if dbuser.Email != "" {
		response := web.WebResponse{
			Data: "Email already in use",
		}
		helper.Response(w, response, 400)
		return
	}

	user.Password, err = services.GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("error in password hash")
	}

	//insert user details in database
	app.Instance.Create(&user)
	response := web.WebResponse{
		Data: user,
	}
	helper.Response(w, response, 200)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()

	var authRequest web.LoginRequest

	helper.ReadFromRequestBody(r, &authRequest)

	err := validate.Struct(authRequest)

	errors := helper.FormatValidationError(err)

	if errors != nil {
		response := web.WebResponse{
			Data: errors,
		}
		helper.Response(w, response, 400)
		return
	}

	var authUser domain.User

	app.Instance.Where("email = ?", authRequest.Email).First(&authUser)

	if authUser.Email == "" {
		response := web.WebResponse{
			Data: "Username or Password is incorrect",
		}
		helper.Response(w, response, 400)
		return
	}

	check := services.CheckPasswordHash(authRequest.Password, authUser.Password)

	if !check {
		response := web.WebResponse{
			Data: "Username or Password is incorrect",
		}
		helper.Response(w, response, 400)
		return
	}

	validToken, err := services.GenerateJwt(authUser.Email, authUser.Role)

	if err != nil {
		response := web.WebResponse{
			Data: "Failed to generate token",
		}
		helper.Response(w, response, 400)
		return
	}

	var token domain.Token
	token.Email = authUser.Email
	token.Role = authUser.Role
	token.TokenString = validToken

	response := web.WebResponse{
		Data: token,
	}
	helper.Response(w, response, 200)
}
