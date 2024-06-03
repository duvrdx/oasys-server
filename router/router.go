package router

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the routes for the application.
func SetupRoutes() *gin.Engine {
	router := gin.Default()
	SetupUser(router)
	SetupGroups(router)
	return router
}
