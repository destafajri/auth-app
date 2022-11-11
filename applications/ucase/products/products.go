package products

import (
	"net/http"

	"github.com/destafajri/auth-app/applications/repository"
	"github.com/gin-gonic/gin"
)

type welcomeHandler struct{
	root repository.UserInterface
}

func NewWelcomeMessage(root repository.UserInterface) *welcomeHandler{
	return &welcomeHandler{
		root,
	}
}

func (handler *welcomeHandler)WelcomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name" : "Desta",
		"status" : "Welcome to My API with Golang-Gin Library",
	})
}