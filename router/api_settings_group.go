package router

import "gvb_server/api"

func (ag ApiGroup) InitSettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	ag.GET("", settingsApi.SettingsApiView)
}
