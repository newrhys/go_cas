package request

type PageInfo struct {
	CurrentPage  		int 		`form:"current_page"`
	PageSize			int 		`form:"page_size"`
}

type ReqID struct {
	ID 					uint64 		`json:"id" binding:"required" label:"ID"`
}
