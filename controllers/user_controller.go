package controllers

import (
	"go-jwt/app"
	"go-jwt/helper"
	"go-jwt/models/domain"
	"go-jwt/models/web"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	var data []domain.User
	app.Instance.Model(&domain.User{}).Scan(&data)

	response := web.WebResponse{
		Data: data,
	}

	helper.Response(w, response, 200)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()

	var user web.User

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
			Data: "User Not Found",
		}
		helper.Response(w, response, 400)
		return
	}

	user.Password, err = GeneratehashPassword(user.Password)
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

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
