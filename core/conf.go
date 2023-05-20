package core

import (
	"fmt"
	"gvb_server/config"
	"gvb_server/global"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// 初始化配置文件
func InitConfig() {
	//yaml文件路径
	const CONFIGFILE = "setting.yaml"
	//配置文件结构体
	c := &config.Config{}
	//读取yaml的byte[]
	yamlConf, err := ioutil.ReadFile(CONFIGFILE)
	if err != nil {
		panic(fmt.Errorf("获取yaml配置错误：%s", err))
	}
	//将yaml文件转化为结构体
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("yaml文件初始化成功")
	//赋值给全局配置
	global.Config = c
}
