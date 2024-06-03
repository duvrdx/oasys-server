package router

import (
	"github.com/duvrdx/oasys-server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUser(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:id", controllers.GetUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
}
