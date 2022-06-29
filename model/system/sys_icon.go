package system

import "wave-admin/model"

type SysIcon struct {
	model.GnModel
	ParentId 		uint64      `gorm:"default:0;comment:父id" json:"parent_id"`
	Name       		string     	`gorm:"comment:名称" json:"name"`
	Children   		[]SysIcon 	`gorm:"-" json:"children"`
}
