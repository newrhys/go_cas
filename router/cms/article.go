package cms

import (
	"github.com/gin-gonic/gin"
	v1 "wave-admin/controller/v1"
	"wave-admin/middleware"
)

type ArticleRouter struct{}

func (s *TagRouter) InitArticleRouter(Router *gin.RouterGroup)  {
	articleRouterWithoutRecord := Router.Group("/article")
	articleRouter := Router.Group("/article").Use(middleware.Record())
	articleApi := v1.ApiGroupApp.CmsApiGroup.ArticleApi
	{
		articleRouter.POST("/article", articleApi.AddArticle)
		articleRouter.PUT("/article/:id", articleApi.UpdateArticle)
		articleRouter.DELETE("/article/:id", articleApi.DeleteArticle)
	}
	{
		articleRouterWithoutRecord.GET("/article/:id", articleApi.GetArticle)
		articleRouterWithoutRecord.GET("/getArticleList", articleApi.ArticleList)
	}
}
