package cms

import "wave-admin/model"

type CmsCategory struct {
	model.GnModel
	Status 			uint8     	`gorm:"default:1;comment:状态：1-正常 2-停用" json:"status"`
	ParentId 		uint64      `gorm:"index;default:0;comment:父id" json:"parent_id"`
	Name 			string      `gorm:"index;comment:分类名" json:"name"`
	Sort 			int         `gorm:"default:0;comment:排序标记" json:"sort"`
}
