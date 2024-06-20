package router

import (
	"gitee.com/trustChain/framework/core/global"
	"gitee.com/trustChain/framework/core/protocol"
	"github.com/gin-gonic/gin"
	"github.com/provider-go/sso"
	"github.com/spf13/viper"
)

func ImportPlugin(r *gin.Engine) {
	publicGroup := r.Group(viper.GetString("service.pre"))
	PluginInit(publicGroup,
		sso.CreatePluginAndDB(global.PluginConfig),
	)
}

func PluginInit(group *gin.RouterGroup, Plugin ...protocol.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}
