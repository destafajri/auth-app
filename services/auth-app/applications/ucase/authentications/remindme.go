package authentications

import (
	"net/http"

	"github.com/destafajri/auth-app/applications/repository"
	"github.com/gin-gonic/gin"
)

type remindUser struct {
	remind repository.UserInterface
}

func NewRemindUser(remind repository.UserInterface) *remindUser {
	return &remindUser{
		remind,
	}
}

func(r remindUser) Remindme(c *gin.Context) {
	var payload struct{
		Phone	string	`json:"phone"`
	}

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors" : err,
		})
		return
	}

	user, err := r.remind.Remindme(payload.Phone)
	if err!= nil {
		c.JSON(http.StatusNotFound, gin.H{
			"errors"	: err,
			"msg"		: "User not Found",
		})
		return
	}

	userRes := convertToAuthResponse(*user)
	c.JSON(http.StatusAccepted, gin.H{
		"data" : userRes,
	})
}