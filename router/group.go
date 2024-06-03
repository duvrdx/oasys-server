package router

import (
	"github.com/duvrdx/oasys-server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupGroups(router *gin.Engine) {
	router.GET("/groups", controllers.GetGroups)
	router.POST("/groups", controllers.CreateGroup)
	router.GET("/groups/:id", controllers.GetGroup)
	router.PUT("/groups/:id", controllers.UpdateGroup)
	router.DELETE("/groups/:id", controllers.DeleteGroup)
	router.POST("/groups/:id/addUser/:user_id", controllers.AddUserToGroup)
	router.POST("/groups/:id/removeUser/:user_id", controllers.RemoveUserFromGroup)
}
