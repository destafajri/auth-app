package agregate

import (
	"net/http"

	"github.com/destafajri/fetch-app/applications/middlewares"
	"github.com/destafajri/fetch-app/applications/helper"
	"github.com/destafajri/fetch-app/applications/repository"
	"github.com/gin-gonic/gin"
)

type getAgregateHandler struct{
	root repository.FetchInterface
}

func NewGetAgregateData(root repository.FetchInterface) *getAgregateHandler{
	return &getAgregateHandler{
		root,
	}
}

func (handler *getAgregateHandler)AgregateDataHandler(c *gin.Context) {
	// middleware
	if middlewares.ROLE != "admin"{
		c.JSON(http.StatusUnauthorized, gin.H{
			"messsage": "Unauthorized as Admin",
		})
		return
	}

	response := helper.Resource()
	
	c.JSON(http.StatusOK, response)
}