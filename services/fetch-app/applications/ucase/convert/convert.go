package convert

import (
	"net/http"

	"github.com/destafajri/fetch-app/applications/middlewares"
	"github.com/destafajri/fetch-app/applications/helper"
	"github.com/destafajri/fetch-app/applications/repository"
	"github.com/gin-gonic/gin"
)

type getCurrencyHandler struct {
	root repository.FetchInterface
}

func NewGetFetchData(root repository.FetchInterface) *getCurrencyHandler {
	return &getCurrencyHandler{
		root,
	}
}

func (handler *getCurrencyHandler) CurrencyDataHandler(c *gin.Context) {
	var payload struct{
		Price string `json:"price"`
	}

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors" : err,
		})
		return
	}

	// middleware
	if middlewares.ROLE == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"messsage": "Unauthorized as member",
		})
		return
	}

	response, err := helper.ConvertCurrency(payload.Price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errro": err,
		})
		return
	}

	c.JSON(http.StatusOK, response)
}