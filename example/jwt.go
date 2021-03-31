package example

import (
	"fmt"
	"time"

	"github.com/quangdangfit/gocommon/jwt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func JWT() {
	jAuth := jwt.New(jwt.WithTokenSecretKey("my-secret-key"), jwt.WithExpiredTime(1*time.Minute))

	user := User{
		ID:       1,
		Username: "username",
		Password: "password",
	}

	token, expTime := jAuth.GenerateToken(user)
	fmt.Printf("Token: %s | Expiried Time: %s", token, expTime)

	data, err := jAuth.ValidateToken(token)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
}
