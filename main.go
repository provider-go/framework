package main

import (
	"gitee.com/trustChain/framework/config"
	"gitee.com/trustChain/framework/core/global"
	"gitee.com/trustChain/framework/router"
)

func main() {

	// 加载配置
	config.Start()
	// 初始化
	global.Initialize()

	// 启动服务并注册
	router.StartRouter()
}
