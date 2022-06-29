package system

import (
	"github.com/gin-gonic/gin"
	v1 "wave-admin/controller/v1"
	"wave-admin/middleware"
)

type UserRouter struct{}

func (s *UserRouter) InitAuthRouter(Router *gin.RouterGroup) {
	baseRouterWithoutRecord := Router.Group("/auth")
	baseRouter := Router.Group("/auth").Use(middleware.Record())
	userApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		baseRouter.POST("/login", userApi.Login)
	}
	{
		baseRouterWithoutRecord.GET("/captcha", userApi.Captcha)
	}
}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouterWithoutRecord := Router.Group("/user")
	userRouter := Router.Group("/user").Use(middleware.Record())
	userApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.POST("/user", userApi.AddUser)
		userRouter.PUT("/user/:id", userApi.UpdateUser)
		userRouter.DELETE("/user/:id", userApi.DeleteUser)
		userRouter.POST("/loginOut", userApi.LoginOut)
		userRouter.POST("/assignRole", userApi.AssignRole)
		userRouter.POST("/changePassword", userApi.ChangePassword)
	}
	{
		userRouterWithoutRecord.GET("/getUserList", userApi.GetUserList)
		userRouterWithoutRecord.GET("/getRoleIdByUserId/:id", userApi.GetRoleIdByUserId)
		userRouterWithoutRecord.GET("/info", userApi.GetUserInfo)
		userRouterWithoutRecord.GET("/getRouteMenuList", userApi.GetRouteMenuList)
	}
}