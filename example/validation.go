package example

import (
	"fmt"

	"github.com/quangdangfit/gocommon/validation"
)

type LoginBody struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Validation() {
	var validator = validation.New()
	var body = LoginBody{
		Email:    "email",
		Password: "password",
	}

	if err := validator.ValidateStruct(body); err != nil {
		fmt.Println(err)
	}
}
