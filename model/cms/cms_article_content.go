package cms

import "wave-admin/model"

type CmsArticleContent struct {
	model.GnModel
	ArticleId 		uint64 		`gorm:"index;comment:文章id" json:"article_id"`
	Description    	string		`gorm:"comment:描述" json:"description"`
	Content    		string		`gorm:"type:text;comment:内容" json:"content"`
}
