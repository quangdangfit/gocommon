package jwt

import (
	"time"
)

// IJWTAuth interface
type IJWTAuth interface {
	GenerateToken(data interface{}) (string, *time.Time)
	ValidateToken(jwtToken string) (map[string]interface{}, error)
}
