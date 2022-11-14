package fetch

import (
	"net/http"

	"github.com/destafajri/fetch-app/applications/helper"
	"github.com/destafajri/fetch-app/applications/repository"
	"github.com/gin-gonic/gin"
)

type getFetchHandler struct{
	root repository.FetchInterface
}

func NewGetFetchData(root repository.FetchInterface) *getFetchHandler{
	return &getFetchHandler{
		root,
	}
}

func (handler *getFetchHandler)FetchDataHandler(c *gin.Context) {
	//middleware
	// if helpermiddleware.ROLE == "admin"{
	// 	return
	// }

	response := helper.Resource()

	c.JSON(http.StatusOK, response)
}