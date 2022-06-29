package system

import (
	"wave-admin/model"
	"time"
	)


type SysRecord struct {
	model.GnModel
	Description		string        `gorm:"comment:描述" json:"description"`			// 描述
	Method       	string        `gorm:"comment:请求方法" json:"method"`				// 请求方法
	Path         	string        `gorm:"comment:请求路径" json:"path"`				// 请求路径
	Status       	int           `gorm:"comment:请求状态" json:"status"`				// 请求状态
	Latency      	time.Duration `gorm:"comment:延迟" json:"latency"` 				// 延迟
	ErrorMessage 	string        `gorm:"comment:错误信息" json:"error_message"`  	// 错误信息
	Ip           	string        `gorm:"comment:请求ip" json:"ip"`					// 请求ip
	Body         	string        `gorm:"type:text;comment:请求Body" json:"body"`	// 请求Body
	Resp         	string        `gorm:"type:text;comment:响应Body" json:"resp"`	// 响应Body
	UserID       	uint64        `gorm:"comment:用户id" json:"user_id"` 			// 用户id
	User         	SysUser       `json:"user"`
}
