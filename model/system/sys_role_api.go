package system

import "wave-admin/model"

type SysRoleApi struct {
	model.GnModel
	RoleId 			uint64			`gorm:"comment:角色ID" json:"role_id"`
	ApiId      		uint64			`gorm:"comment:api ID" json:"api_id"`
}
