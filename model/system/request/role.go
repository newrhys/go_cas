package request

import (
	"wave-admin/model/common/request"
)

type ReqRole struct {
	ID 					uint64 		`json:"id"`
	RoleName			string		`json:"role_name" binding:"required,max=255,min=2" label:"角色名称"`
	Remark  			string 		`json:"remark"`
	ParentId			uint64 		`json:"parent_id"`
	Status 				uint8		`json:"status"`
}

type RoleList struct {
	RoleName			string		`form:"role_name"`
	Status 				string		`form:"status"`
	request.PageInfo
}

type AssignSave struct {
	ID 					uint64 		`json:"id"`
	MenuList 			[]int 		`json:"menu_list"`
	ApiList 			[]int 		`json:"api_list"`
}

