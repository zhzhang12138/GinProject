package initialize

import (
	_ "gin-project/docs"
	"gin-project/global"
	"gin-project/middleware"
	"gin-project/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RouterGroupApp.System
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")

	// 公共接口-无需鉴权
	PublicGroup := Router.Group("")
	// 健康监测
	{
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		systemRouter.InitUserRouter(PublicGroup) // 注册角色路由
	}

	// 私有接口-需鉴权
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		//systemRouter.InitAuthorityRouter(PrivateGroup) // 注册角色路由
	}

	global.GVA_LOG.Info("router register success")
	return Router
}
