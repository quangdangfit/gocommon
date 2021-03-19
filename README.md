# gocommon

Golang common libraries.

## Installation

`go get -u github.com/quangdangfit/gocommon`

Note that zap only supports the two most recent minor versions of Go.

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

## Contributing

If you want to contribute to this boilerplate, clone the repository and just
start making pull requests.