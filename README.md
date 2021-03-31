# Go Common Libraries

[![Build Status](https://travis-ci.org/quangdangfit/gocommon.svg?branch=master)](https://travis-ci.org/quangdangfit/gocommon)
[![codecov](https://codecov.io/gh/quangdangfit/gocommon/branch/master/graph/badge.svg)](https://codecov.io/gh/quangdangfit/gocommon)
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/quangdangfit/gocommon)

## Quick Start
#### Logger:
```go
package main
import (
    "github.com/quangdangfit/gocommon/logger"
)
func main() {
    logger.Initialize("production")
    logger.Info("This info log")
}
```
#### Validation:
```go
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
```
#### Redis:
```go
package main
import (
	"fmt"
	
	"github.com/quangdangfit/gocommon/redis"
)
func main() {
	var conf = redis.Config{
		Address:  "localhost:6379",
		Password: "password",
		Database: 1,
	}
	var r = redis.New(conf)
	if err := r.Set("key", "value"); err != nil {
		fmt.Println(err)
	}
	var value string
	if err := r.Get("key", &value); err != nil {
		fmt.Println(err)
	}
	fmt.Println("value: ", value)
}
```
#### JWT
```go
package main
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
func main() {
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
```
#### Solr
```go
package main
import (
	"encoding/json"
	"fmt"
	
	"github.com/quangdangfit/gocommon/solr"
)
func main() {
	var conf = solr.Config{
		URL:      "http://localhost:8983/",
		Core:     "core",
		User:     "user",
		Password: "password",
	}
	s, err := solr.New(conf)
	if err != nil {
		fmt.Println(err)
	}
	data := map[string]interface{}{
		"key": "value",
	}
	b, err := json.Marshal(&data)
	if err != nil {
		fmt.Println(err)
	}
	if err := s.Add(b); err != nil {
		fmt.Println(err)
	}
}
```
## Contributing
If you want to contribute to this boilerplate, clone the repository and just
start making pull requests.