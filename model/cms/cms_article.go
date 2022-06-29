package cms

import (
	"wave-admin/model"
	"wave-admin/model/system"
)

type CmsArticle struct {
	model.GnModel
	Status 				uint8     			`gorm:"default:0;comment:状态 0-待审核，1-发布，2-删除" json:"status"`
	Title       		string      		`gorm:"index;comment:文章标题" json:"title"`
	UserID 				uint64				`gorm:"index;comment:创建者ID" json:"user_id"`
	EditId				uint64				`gorm:"index;comment:编辑者ID" json:"edit_id"`
	Img 				string 				`gorm:"comment:封面图" json:"img"`
	KeyWord				string 				`gorm:"type:varchar(50);comment:短标题" json:"key_word"`
	Type 				uint8 				`gorm:"index;default:0;comment:文章类型 0-文字，1-图片，2-视频" json:"type"`
	CategoryId 			uint64 				`gorm:"index;comment:分类id" json:"category_id"`
	User         		system.SysUser      `json:"user"`
}
