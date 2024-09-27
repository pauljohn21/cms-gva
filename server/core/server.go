package core

import (
	"fmt"

	"github.com/pauljohn21/cms-gva/server/global"
	"github.com/pauljohn21/cms-gva/server/initialize"
	"github.com/pauljohn21/cms-gva/server/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)

	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 cms-gva
	当前版本:v0.0.1
	插件市场:https://plugin.gin-vue-admin.com
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080`, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
