package request

import "wave-admin/model/common/request"

type ReqArticle struct {
	Title       	string      `json:"title" binding:"required,max=255,min=2" label:"文章标题"`
	Img 			string 		`json:"img" binding:"required" label:"文章图片"`
	KeyWord			string 		`json:"key_word"`
	CategoryId 		uint64 		`json:"category_id" binding:"required" label:"文章分类"`
	Status 			uint8     	`json:"status" binding:"required" label:"状态"`
	TagIds 			string		`json:"tag_ids"`
	Description    	string		`json:"description" binding:"required" label:"描述"`
	Content    		string		`json:"content" binding:"required" label:"文章内容"`
}

type ArticleList struct {
	Title       	string 		`form:"title"`
	Status 			string		`form:"status"`
	request.PageInfo
}
