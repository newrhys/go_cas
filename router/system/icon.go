package system

import (
	"github.com/gin-gonic/gin"
	v1 "wave-admin/controller/v1"
)

type IconRouter struct{}

func (s *IconRouter) InitIconRouter(Router *gin.RouterGroup) {
	iconRouter := Router.Group("/icon")
	iconApi := v1.ApiGroupApp.SystemApiGroup.IconApi
	iconRouter.GET("/getIconList", iconApi.IconList)
}