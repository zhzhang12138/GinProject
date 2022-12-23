package main

import (
	"gin-project/core"
	"gin-project/global"
	"gin-project/initialize"
	"go.uber.org/zap"
)

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample Server pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	// 1、初始化Viper
	global.GVA_VP = core.Viper()
	// 2、其他初始化
	initialize.OtherInit()
	// 3、初始化zap日志库
	global.GVA_LOG = core.Zap()
	// 4、替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(global.GVA_LOG)
	// 5、gorm连接数据库
	global.GVA_DB = initialize.Gorm()
	// 6、多数据库支持
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	// 7、启动项目
	core.RunServer()

}
