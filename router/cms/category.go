package cms

import (
	"github.com/gin-gonic/gin"
	v1 "wave-admin/controller/v1"
	"wave-admin/middleware"
)

type CategoryRouter struct{}

func (s *CategoryRouter) InitCategoryRouter(Router *gin.RouterGroup)  {
	categoryRouterWithoutRecord := Router.Group("/category")
	categoryRouter := Router.Group("/category").Use(middleware.Record())
	categoryApi := v1.ApiGroupApp.CmsApiGroup.CategoryApi
	{
		categoryRouter.POST("/category", categoryApi.AddCategory)
		categoryRouter.PUT("/category/:id", categoryApi.UpdateCategory)
		categoryRouter.DELETE("/category/:id", categoryApi.DeleteCategory)
	}
	{
		categoryRouterWithoutRecord.GET("/getCategoryList", categoryApi.CategoryList)
		categoryRouterWithoutRecord.GET("/parent", categoryApi.ParentList)
		categoryRouterWithoutRecord.GET("/getSort/:id", categoryApi.GetSort)
	}
}