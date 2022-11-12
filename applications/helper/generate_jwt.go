package helper

import (
	"fmt"
	"time"

	helpermiddleware "github.com/destafajri/auth-app/applications/middlewares/helperMiddleware"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwtToken(name, phone, role string) (string, error) {
	//proses pembuatan token jwt
	expTime := time.Now().Add(time.Minute * 300)
	claims := helpermiddleware.JWTClaim{
		Name:   name,
		Phone: phone,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "github.com/destafajri",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	//medeklarasikan algoritma yang akan digunakan untuk signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//signed token
	token, err := tokenAlgo.SignedString(helpermiddleware.JWT_SECRET_KEY)
	if err != nil {
		fmt.Println("signed token error")
		return "", err
	}

	return token, nil
}