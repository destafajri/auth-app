package authentications

import (
	"net/http"
	_ "strconv"

	"github.com/destafajri/auth-app/applications/entity"
	"github.com/destafajri/auth-app/applications/repository"
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
	//uuid
	id := uuid.New()
	//generate password

	//payload untuk masuk query
	register := entity.UserEntity{
		ID		: id.String(),
		Name 	: payload.Name,
		Phone 	: payload.Phone,
		Role 	: payload.Role,
		Password: "ongoing",
	}

	err = r.register.Register(&register)
	if err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors"	: err,
			"msg"		: "query",
		})
		return
	}

	regisRes := convertToAuthResponse(register)
	c.JSON(http.StatusAccepted, gin.H{
		"data" : regisRes,
	})
}