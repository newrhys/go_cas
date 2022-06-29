package system

import (
	"github.com/gin-gonic/gin"
	v1 "wave-admin/controller/v1"
	"wave-admin/middleware"
)

type RoleRouter struct{}

func (s *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	roleRouterWithoutRecord := Router.Group("/role")
	roleRouter := Router.Group("/role").Use(middleware.Record())
	roleApi := v1.ApiGroupApp.SystemApiGroup.RoleApi
	{
		roleRouter.POST("/role", roleApi.AddRole)
		roleRouter.PUT("/role/:id", roleApi.UpdateRole)
		roleRouter.DELETE("/role/:id", roleApi.DeleteRole)
		roleRouter.POST("/roleAssignSave", roleApi.AssignSave)
	}
	{
		roleRouterWithoutRecord.GET("/getRoleList", roleApi.RoleList)
		roleRouterWithoutRecord.GET("/getAssignPermissionTree/:id", roleApi.GetAssignPermissionTree)
		roleRouterWithoutRecord.GET("/getAssignPermissionApi/:id", roleApi.GetAssignPermissionApi)
	}
}
