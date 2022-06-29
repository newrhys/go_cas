package system

import "wave-admin/model"

type SysApi struct {
	model.GnModel
	ParentId 		uint64      `gorm:"default:0;comment:父id" json:"parent_id"`
	Description    	string		`gorm:"comment:描述" json:"description"`
	Path 			string      `gorm:"comment:api路径" json:"path"`
	Method 			string 		`gorm:"comment:方法" json:"method"`
	Children   		[]SysApi 	`gorm:"-" json:"children"`
}
