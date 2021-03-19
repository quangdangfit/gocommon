package main

import (
	"fmt"

	"github.com/quangdangfit/gocommon/validation"
)

type LoginBody struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func main() {
	var validator = validation.New()
	var body = LoginBody{
		Email:    "email",
		Password: "password",
	}

	if err := validator.ValidateStruct(body); err != nil {
		fmt.Println(err)
	}
}
