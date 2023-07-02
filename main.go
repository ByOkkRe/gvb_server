package main

import (
	"gvb_server/core"
	"gvb_server/flag"
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

	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	router := router.InitRouter()
	global.Log.Infof("系统运行在[%s]", global.Config.System.Addr())
	err := router.Run(global.Config.System.Addr())
	if err != nil {
		global.Log.Fatal(err.Error())
	}
}
