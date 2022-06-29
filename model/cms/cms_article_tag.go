package cms

import "wave-admin/model"

type CmsArticleTag struct {
	model.GnModel
	ArticleId 		uint64 		`gorm:"index;comment:文章id" json:"article_id"`
	TagId 			uint64 		`gorm:"index;comment:标签id" json:"tag_id"`
}
