package request

import "wave-admin/model/common/request"

type ReqApi struct {
	ParentId 		uint64 		`json:"parent_id"`
	Description    	string		`json:"description" binding:"required" label:"描述"`
	Path 			string      `json:"path"`
	Method 			string 		`json:"method"`
}

type ApiList struct {
	Description    	string		`form:"description"`
	Path 			string      `form:"path"`
	Method 			string 		`form:"method"`
	request.PageInfo
}

type AssignApi struct {
	ID 					uint64 		`json:"id"`
	List 				[]int 		`json:"list"`
}