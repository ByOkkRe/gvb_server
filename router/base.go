package router

import (
	"gvb_server/global"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	apiGroup := ApiGroup{}
	apiGroup.InitApiRouterGroup(router)
	return router
}
