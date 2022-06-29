package cms

import "wave-admin/service"

type ApiGroup struct {
	CategoryApi
	TagApi
	ArticleApi
}

var (
	categoryService = service.ServiceGroupApp.CmsServiceGroup.CategoryService
	tagService = service.ServiceGroupApp.CmsServiceGroup.TagService
	articleService = service.ServiceGroupApp.CmsServiceGroup.ArticleService
)
