package request

type ReqMenu struct {
	ParentId 			uint64 			`json:"parent_id"`
	Type				uint8			`json:"type"`
	Path 				string 			`json:"path"`
	Name 				string 			`json:"name"`
	Hidden 				int 			`json:"hidden"`
	Component 			string 			`json:"component"`
	Sort 				int 			`json:"sort" binding:"required" label:"排序"`
	Icon 				string 			`json:"icon"`
	Code				string			`json:"code" binding:"required,max=255,min=2" label:"权限code"`
	KeepAlive   		int   			`json:"keep_alive"`
	Title       		string 			`json:"title" binding:"required,max=255,min=2" label:"菜单名称"`
	IsLink				int 			`json:"is_link"`
	Status 				uint8			`json:"status" binding:"required"`
}

type MenuList struct {
	Title       		string 			`form:"title"`
	Status 				string			`form:"status"`
}