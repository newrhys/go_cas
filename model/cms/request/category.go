package request

type ReqCategory struct {
	ParentId 		uint64 		`json:"parent_id"`
	Status 			uint8     	`json:"status" binding:"required" label:"状态"`
	Name 			string      `json:"name" binding:"required" label:"分类名称"`
	Sort 			int 		`json:"sort" binding:"required" label:"排序"`
}

type CategoryList struct {
	Name       		string 		`form:"name"`
	Status 			string		`form:"status"`
}