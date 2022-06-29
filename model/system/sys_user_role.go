package system

import "wave-admin/model"

type SysUserRole struct {
	model.GnModel
	UserId			uint64			`gorm:"comment:管理员ID" json:"user_id"`
	RoleId 			uint64			`gorm:"comment:角色ID" json:"role_id"`
}