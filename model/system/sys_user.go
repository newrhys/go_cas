package system

import (
	uuid "github.com/satori/go.uuid"
	"wave-admin/model"
)

type SysUser struct {
	model.GnModel
	UUID 		uuid.UUID 		`gorm:"comment:用户UUID" json:"uuid"`
	Username 	string 			`gorm:"index;comment:用户名" json:"username"`
	Password 	string 			`gorm:"type:varchar(32);not null;comment:密码" json:"password"`
	Nickname 	string 			`gorm:"comment:昵称" json:"nickname"`
	Salt 		string 			`gorm:"type:varchar(10);not null;comment:密码加密串" json:"salt"`
	Avatar 		string 			`gorm:"comment:头像" json:"avatar"`
	Status 		uint8 			`gorm:"default:1;comment:状态：1-正常 2-停用" json:"status"`
	Mobile 		string 			`gorm:"type:varchar(11);not null;comment:手机号" json:"mobile"`
	Email 		string 			`gorm:"comment:邮箱" json:"email"`
}