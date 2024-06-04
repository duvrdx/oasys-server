package router

import (
	"fmt"

	"github.com/duvrdx/oasys-server/config"
	"github.com/duvrdx/oasys-server/controllers"
	"github.com/duvrdx/oasys-server/services"

	"github.com/gin-gonic/gin"
)

func SetupUser(router *gin.Engine) {
	fmt.Println(config.DB)
	userService := services.NewUserService(config.DB)
	userController := controllers.NewUserController(userService)

	router.GET("/users", userController.GetUsers)
	router.POST("/users", userController.CreateUser)
	router.GET("/users/:id", userController.GetUser)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)
}
