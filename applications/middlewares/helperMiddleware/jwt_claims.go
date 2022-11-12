package helpermiddleware

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

// helper variable
var JWT_SECRET_KEY = []byte(os.Getenv("KEY_JWT"))
var ROLE string
var CLAIMS JWTClaim

// claims struct
type JWTClaim struct {
	Name	string
	Phone	string
	Role	string
	jwt.RegisteredClaims
}