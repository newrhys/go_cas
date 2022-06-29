package cms

import "wave-admin/model"

type CmsTag struct {
	model.GnModel
	Status 			uint8     	`gorm:"default:1;comment:状态：1-正常 2-停用" json:"status"`
	Name 			string      `gorm:"index;comment:标签名" json:"name"`
	Sort 			int         `gorm:"default:0;comment:排序标记" json:"sort"`
}
