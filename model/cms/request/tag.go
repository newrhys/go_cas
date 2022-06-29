package request

import "wave-admin/model/common/request"

type ReqTag struct {
	Status 			uint8     	`json:"status" binding:"required" label:"状态"`
	Name 			string      `json:"name" binding:"required" label:"标签名称"`
	Sort 			int 		`json:"sort" binding:"required" label:"排序"`
}

type TagList struct {
	Name       		string 		`form:"name"`
	Status 			string		`form:"status"`
	request.PageInfo
}