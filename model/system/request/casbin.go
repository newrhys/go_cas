package request

// Casbin info structure
type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

func DefaultCasbin() []CasbinInfo {
	return []CasbinInfo{
		{Path: "/api/v1/user/loginOut", Method: "POST"},
		{Path: "/api/v1/user/info", Method: "GET"},
		{Path: "/api/v1/user/getRouteMenuList", Method: "GET"},
	}
}
