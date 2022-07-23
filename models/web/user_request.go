package web

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `validate:"required,min=3"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
	Role     string `validate:"required"`
}
