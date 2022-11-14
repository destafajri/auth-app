package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

//validate token
func tokenValidation(c *gin.Context, token *jwt.Token, err error) error{
	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		switch v.Errors {
		case jwt.ValidationErrorSignatureInvalid:
			// token invalid
			response := "Unauthorized"
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": response,
			})
			return err
		case jwt.ValidationErrorExpired:
			// token expired
			response := "Unauthorized, Token expired!"
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": response,
			})
			return err
		default:
			response := "Unauthorized"
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": response,
			})
			return err
		}
	}

	if !token.Valid {
		response := "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": response,
		})
		return err
	}

	return nil
}