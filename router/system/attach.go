package system

import (
	"github.com/gin-gonic/gin"
	v1 "wave-admin/controller/v1"
)

type AttachRouter struct{}

func (s *AttachRouter) InitAttachRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("/attach")
	attachApi := v1.ApiGroupApp.SystemApiGroup.AttachApi
	{
		baseRouter.POST("/upload", attachApi.Upload)
	}
}
