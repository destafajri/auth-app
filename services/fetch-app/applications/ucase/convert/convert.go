package convert

import (
	"net/http"

	helpermiddleware "github.com/destafajri/auth-app/applications/middlewares/helperMiddleware"
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
	if helpermiddleware.ROLE == "" {
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