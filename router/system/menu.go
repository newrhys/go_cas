package system

import (
	"github.com/gin-gonic/gin"
	v1 "wave-admin/controller/v1"
	"wave-admin/middleware"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	menuRouterWithoutRecord := Router.Group("/menu")
	menuRouter := Router.Group("/menu").Use(middleware.Record())
	menuApi := v1.ApiGroupApp.SystemApiGroup.MenuApi
	{
		menuRouter.POST("/menu", menuApi.AddMenu)
		menuRouter.PUT("/menu/:id", menuApi.UpdateMenu)
		menuRouter.DELETE("/menu/:id", menuApi.DeleteMenu)
	}
	{
		menuRouterWithoutRecord.GET("/parent", menuApi.ParentList)
		menuRouterWithoutRecord.GET("/getSort/:id", menuApi.GetSort)
		menuRouterWithoutRecord.GET("/getMenuList", menuApi.MenuList)
	}
}
