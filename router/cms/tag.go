package cms

import (
	"github.com/gin-gonic/gin"
	v1 "wave-admin/controller/v1"
	"wave-admin/middleware"
)

type TagRouter struct{}

func (s *TagRouter) InitTagRouter(Router *gin.RouterGroup)  {
	tagRouterWithoutRecord := Router.Group("/tag")
	tagRouter := Router.Group("/tag").Use(middleware.Record())
	tagApi := v1.ApiGroupApp.CmsApiGroup.TagApi
	{
		tagRouter.POST("/tag", tagApi.AddTag)
		tagRouter.PUT("/tag/:id", tagApi.UpdateTag)
		tagRouter.DELETE("/tag/:id", tagApi.DeleteTag)
	}
	{
		tagRouterWithoutRecord.GET("/getTagList", tagApi.TagList)
		tagRouterWithoutRecord.GET("/getSort", tagApi.GetSort)
		tagRouterWithoutRecord.GET("/getSelectTagList", tagApi.SelectTagList)
	}
}
