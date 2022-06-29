package system

import (
	"github.com/gin-gonic/gin"
	v1 "wave-admin/controller/v1"
)

type RecordRouter struct{}

func (s *RoleRouter) InitRecordRouter(Router *gin.RouterGroup) {
	recordRouter := Router.Group("/record")
	recordApi := v1.ApiGroupApp.SystemApiGroup.RecordApi
	recordRouter.GET("/getRecordList", recordApi.RecordList)
	recordRouter.DELETE("/record/:id", recordApi.DeleteRecord)
}