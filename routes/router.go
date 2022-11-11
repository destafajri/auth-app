package routes

import (
	"github.com/destafajri/auth-app/applications/repository"
	"github.com/destafajri/auth-app/applications/ucase/authentications"
	"github.com/destafajri/auth-app/applications/ucase/products"
	"github.com/destafajri/auth-app/db"
	"github.com/gin-gonic/gin"
)

func Router() {
	db:= db.DBcon()

	//Repository
	userRepository := repository.NewUser(db)
	//Ucase
	register := authentications.NewRegisterUser(userRepository)
	product := products.NewWelcomeMessage(userRepository)

	//router default setting
	router := gin.Default()
	//versioning api
	api := router.Group("/api").Use()

	//router path
	router.POST("/register", register.RegisterHandler )
	//api path for root request
	api.GET("/products", product.WelcomeHandler)
	
	router.Run(":9000")
}