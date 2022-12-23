package system

import (
	v1 "gin-project/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("admin_register", baseApi.Register) // 注册-管理员账号
		userRouter.POST("getUserList", baseApi.GetUserList) // 分页获取-用户列表
		userRouter.PUT("setUserInfo", baseApi.SetUserInfo)  // 设置-用户信息
		userRouter.DELETE("deleteUser", baseApi.DeleteUser) // 删除用户
	}

}
