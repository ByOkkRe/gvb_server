package router

import "github.com/gin-gonic/gin"

type ApiGroup struct {
	*gin.RouterGroup
}

func (ApiGroup) InitApiRouterGroup(router *gin.Engine) {
	apiRG := router.Group("/api")
	{
		settingsRG := apiRG.Group("/settings")
		{
			var apiGroup = ApiGroup{settingsRG}
			apiGroup.InitSettingsRouter()
		}

	}
}
