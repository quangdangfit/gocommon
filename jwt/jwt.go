package jwt

import (
	"encoding/json"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Define constants
const (
	UserDataKey = "user_data"
)

// Auth struct
type Auth struct {
	TokenExpiredTime time.Duration
	TokenSecretKey   string
	SigningMethod    string
}

// New Auth object
func New(opts ...Option) IJWTAuth {
	opt := getOption(opts...)
	return &Auth{
		TokenExpiredTime: opt.TokenExpiredTime,
		TokenSecretKey:   opt.TokenSecretKey,
		SigningMethod:    opt.SigningMethod,
	}
}

// GenerateToken generate jwt token
func (a *Auth) GenerateToken(data interface{}) (string, *time.Time) {
	exp := time.Now().Add(a.TokenExpiredTime)
	tokenContent := jwt.MapClaims{
		"payload": map[string]interface{}{
			UserDataKey: data,
		},
		"exp": exp.Unix(),
	}
	jwtToken := jwt.NewWithClaims(
		jwt.GetSigningMethod(a.SigningMethod),
		tokenContent,
	)
	token, err := jwtToken.SignedString([]byte(a.TokenSecretKey))
	if err != nil {
		return "", nil
	}

	return token, &exp
}

// ValidateToken validate jwt token
func (a *Auth) ValidateToken(jwtToken string) (map[string]interface{}, error) {
	tokenData := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(jwtToken, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.TokenSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	b, err := json.Marshal(tokenData["payload"])
	if err != nil {
		return nil, err
	}

	var payload map[string]interface{}
	err = json.Unmarshal(b, &payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
