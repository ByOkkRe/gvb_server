package settings_api

import (
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

func (s SettingsApi) SettingsApiView(c *gin.Context) {
	res.OK(c)
}
