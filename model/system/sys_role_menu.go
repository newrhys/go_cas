package system

import "wave-admin/model"

type SysRoleMenu struct {
	model.GnModel
	RoleId 			uint64			`gorm:"comment:角色ID" json:"role_id"`
	MenuId      	uint64			`gorm:"comment:菜单ID" json:"menu_id"`
}
