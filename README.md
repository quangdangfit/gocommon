# gocommon

Golang common libraries.

## Installation

`go get -u github.com/quangdangfit/gocommon`

Note that zap only supports the two most recent minor versions of Go.

## Quick Start

#### Logger:

```go
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
import (
    "github.com/quangdangfit/gocommon/validation"
)

type LoginBody struct {
    Email    string `json:"email" validate:"required"`
    Password string `json:"password" validate:"required"`
}

func main() {
    var validator = validation.New()
    var body = LoginBody{
        Email: "email",
        Password: "password",
    }
    
    if err := validator.ValidateStruct(body); err != nil {
        return err
    }
}
```

## Development Status: Stable

All APIs are finalized, and no breaking changes will be made in the 1.x series
of releases. Users of semver-aware dependency management systems should pin zap
to `^1`.

## Contributing

If you want to contribute to this boilerplate, clone the repository and just
start making pull requests.