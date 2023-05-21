package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/router"
)

func main() {
	//初始化配置文件
	core.InitConfig()
	//初始化日志
	global.Log = core.Logger()
	//初始化gorm
	global.DB = core.Gorm()

	router := router.InitRouter()
	global.Log.Infof("系统运行在[%s]", global.Config.System.Addr())
	router.Run(global.Config.System.Addr())
}
