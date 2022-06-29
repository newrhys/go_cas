package system

import (
	"github.com/gin-gonic/gin"
	v1 "wave-admin/controller/v1"
	"wave-admin/middleware"
)

type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouterWithoutRecord := Router.Group("/api")
	apiRouter := Router.Group("/api").Use(middleware.Record())
	sysApiApi := v1.ApiGroupApp.SystemApiGroup.SysApiApi
	{
		apiRouter.POST("/api", sysApiApi.AddApi)
		apiRouter.PUT("/api/:id", sysApiApi.UpdateApi)
		apiRouter.DELETE("/api/:id", sysApiApi.DeleteApi)
	}
	{
		apiRouterWithoutRecord.GET("/getApiList", sysApiApi.ApiList)
		apiRouterWithoutRecord.GET("/parent", sysApiApi.ParentList)
	}

}
