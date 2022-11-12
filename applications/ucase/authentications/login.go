package authentications

import (
	"net/http"

	"github.com/destafajri/auth-app/applications/helper"
	"github.com/destafajri/auth-app/applications/repository"
	"github.com/gin-gonic/gin"
)

type loginUser struct {
	login repository.UserInterface
}

func NewLoginUser(login repository.UserInterface) *loginUser {
	return &loginUser{
		login,
	}
}

func(log loginUser) Login(c *gin.Context){
	var payload struct{
		Phone		string `json:"phone"`
		Password	string `json:"password"`
	}

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors" : err,
		})
		return
	}

	us, err := log.login.Login(payload.Phone)
	if err != nil {
		return
	}

	//password check
	if payload.Password != us.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "You're Unauthorized",
		})
		return
	}

	//Generate JWT
	token, err := helper.GenerateJwtToken(us.Name, us.Phone, us.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "generate jwt",
		})		
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}