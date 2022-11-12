package authentications

import (
	"net/http"
	_ "strconv"

	"github.com/destafajri/auth-app/applications/entity"
	"github.com/destafajri/auth-app/applications/helper"
	"github.com/destafajri/auth-app/applications/repository"
	"github.com/destafajri/auth-app/applications/validations"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type registerUser struct{
	register repository.UserInterface
}

func NewRegisterUser(register repository.UserInterface) *registerUser{
	return &registerUser{
		register,
	}
}

func(r registerUser)RegisterHandler(c *gin.Context) {
	var payload struct{
		Name		string	`json:"name"`
		Phone		string	`json:"phone"`
		Role		string	`json:"role"`
	}
	
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors" : err,
		})
		return
	}

	//validation input phone number
	phone := validations.Phonenumber(payload.Phone)
	if !phone {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors" : "Invalid Phone Number Format",
		})
		return
	}

	//generate uuid
	id := uuid.New()

	//generate password
	pass := helper.GeneratePassword(4,1,1,1)

	//payload untuk masuk query
	register := entity.UserEntity{
		ID		: id.String(),
		Name 	: payload.Name,
		Phone 	: payload.Phone,
		Role 	: payload.Role,
		Password: pass,
	}

	err = r.register.Register(&register)
	if err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors"	: err,
			"msg"		: "nomor telah terdaftar",
		})
		return
	}

	regisRes := convertToAuthResponse(register)
	c.JSON(http.StatusAccepted, gin.H{
		"data" : regisRes,
	})
}