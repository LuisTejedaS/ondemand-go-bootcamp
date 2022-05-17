package router

import (
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	return router
}

// CreateRouterGroup creates the `/v1` router group.
func CreateRouterGroup(router *gin.Engine) *gin.RouterGroup {
	v1 := router.Group("/v1")

	return v1
}
