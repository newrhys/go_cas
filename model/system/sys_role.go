package system

import "wave-admin/model"

type SysRole struct {
	model.GnModel
	Status 			uint8 				`gorm:"default:1;comment:状态：1-正常 2-停用" json:"status"`
	RoleName   		string         		`gorm:"not null;comment:角色名" json:"role_name"`
	ParentId        uint64         		`gorm:"not null;default:0;comment:父角色ID" json:"parent_id"`
	Remark			string 				`gorm:"comment:角色备注" json:"remark"`
	DefaultRouter   string         		`gorm:"comment:默认菜单;default:dashboard" json:"default_router"`
}
