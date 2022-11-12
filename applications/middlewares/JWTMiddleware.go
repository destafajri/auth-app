package middlewares

import (
	"strings"

	helpermiddleware "github.com/destafajri/auth-app/applications/middlewares/helperMiddleware"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//mengambil token value
		tokenValue := c.Request.Header.Get("Authorization")

		//parsing authorization bearer
		tokenString := strings.Replace(tokenValue, "Bearer ", "", -1)

		//claims variable
		claims := &helpermiddleware.JWTClaim{}

		// parsing token jwt
		tokenJwt, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return helpermiddleware.JWT_SECRET_KEY, nil
		})

		//cek token validation
		err = helpermiddleware.TokenValidation(c, tokenJwt, err)
		if err != nil {
			return
		}

		// assign role
		helpermiddleware.ROLE = claims.Role
		helpermiddleware.CLAIMS = *claims 
		
		c.Next()
	}
}