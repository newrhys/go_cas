package system

import "wave-admin/model"

type SysMenu struct {
	model.GnModel
	Status 			uint8     	`gorm:"default:1;comment:状态：1-正常 2-停用" json:"status"`
	ParentId 		uint64      `gorm:"default:0;comment:父id" json:"parent_id"`
	Type 			uint8       `gorm:"default:0;comment:菜单类型：0-目录 1-菜单 2-按钮" json:"type"`
	Level   		uint 		`gorm:"default:0;comment:菜单层级" json:"level"`
	Path 			string      `gorm:"comment:路由path" json:"path"`
	Name 			string      `gorm:"comment:路由参数" json:"name"`
	Hidden 			bool      	`gorm:"default:0;comment:是否隐藏：0-否 1-是" json:"hidden"`
	Component 		string     	`gorm:"comment:对应前端文件路径" json:"component"`
	Sort 			int         `gorm:"default:0;comment:排序标记" json:"sort"`
	Icon 			string      `gorm:"comment:菜单图标" json:"icon"`
	Code			string      `gorm:"comment:权限code" json:"code"`
	KeepAlive   	bool        `gorm:"default:1;comment:是否缓存：0-否 1-是" json:"keep_alive"`
	Title       	string      `gorm:"comment:菜单名" json:"title"`
	IsLink      	bool        `gorm:"default:0;comment:是否外链：0-否 1-是" json:"is_link"`
	Children   		[]SysMenu 	`gorm:"-" json:"children"`
}