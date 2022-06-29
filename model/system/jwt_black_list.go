package system

import "wave-admin/model"

type JwtBlacklist struct {
	model.GnModel
	Jwt string `gorm:"type:text;comment:jwt token string" json:"jwt"`
}