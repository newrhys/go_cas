package initialize

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	_ "wave-admin/docs"
	"wave-admin/global"
	"wave-admin/middleware"
	"wave-admin/router"
)

func InitRouter() *gin.Engine {
	if global.GnConfig.System.Mode == gin.TestMode {
		gin.SetMode(gin.TestMode)
	} else if global.GnConfig.System.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	fs := global.GnConfig.Local.Path
	r.StaticFS(fs, http.Dir("./"+fs))

	// 跨域
	r.Use(middleware.Cors()) // 如需跨域可以打开
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GnLog.Info("register swagger handler")

	systemRouter := router.RouterGroupApp.System
	cmsRouter := router.RouterGroupApp.Cms
	PublicGroup := r.Group("/api/v1")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		systemRouter.InitAuthRouter(PublicGroup) 	// 基础功能路由
		systemRouter.InitAttachRouter(PublicGroup) 	// 文件上传
	}

	PrivateGroup := r.Group("/api/v1")
	PrivateGroup.Use(middleware.AuthMiddleware()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitUserRouter(PrivateGroup) 	// 用户
		systemRouter.InitRoleRouter(PrivateGroup) 	// 角色
		systemRouter.InitMenuRouter(PrivateGroup)	// 菜单
		systemRouter.InitApiRouter(PrivateGroup)	// API
		systemRouter.InitIconRouter(PrivateGroup)	// 图标
		systemRouter.InitRecordRouter(PrivateGroup) // 操作记录
		cmsRouter.InitCategoryRouter(PrivateGroup) 	// 文章分类
		cmsRouter.InitTagRouter(PrivateGroup)		// 文章标签
		cmsRouter.InitArticleRouter(PrivateGroup)	// 文章管理
	}

	return r
}