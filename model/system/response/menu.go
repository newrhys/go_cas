package response

import "time"

type ParentMenu struct {
	ID 				uint64     `json:"id"`
	Value 			uint64      `json:"value"`
	Title 			string      `json:"label"`
	Children   		[]ParentMenu `json:"children"`
}

type AssignRoleMenu struct {
	CheckList 		[]uint64 		`json:"check_list"`
	MenuList 		interface{}		`json:"menu_list"`
}

type AssignRoleApi struct {
	CheckList 		[]uint64 		`json:"check_list"`
	ApiList 		interface{}		`json:"api_list"`
}

type MenuList struct {
	ID        		uint64 			`json:"id"`
	CreatedAt 		time.Time 		`json:"created_at"`
	Status 			uint8 			`json:"status"`
	ParentId 		uint64 			`json:"parent_id"`
	Type 			uint8			`json:"type"`
	Path 			string     		`json:"path"`
	Name 			string     		`json:"name"`
	Hidden 			bool     		`json:"hidden"`
	Component 		string    		`json:"component"`
	Sort 			int        		`json:"sort"`
	Icon 			string     		`json:"icon"`
	Code			string     		`json:"code"`
	KeepAlive   	bool          	`json:"keep_alive"`
	Title       	string        	`json:"title"`
	IsLink      	bool         	`json:"is_link"`
	Children   		[]MenuList 		`json:"children"`
}

type MenuMeta struct {
	Icon 			string 			`json:"icon"`
	Title       	string 			`json:"title"`
	Code			string			`json:"code"`
	KeepAlive   	bool          	`json:"keep_alive"`
}

type TreeList struct {
	Id 				uint64 			`json:"id"`
	ParentId 		uint64 			`json:"parent_id"`
	Type 			uint8       	`json:"type"`
	Name      		string     		`json:"name"`
	Path      		string     		`json:"path"`
	Component 		string     		`json:"component"`
	Hidden    		bool      	 	`json:"hidden"`
	IsLink      	bool 			`json:"is_link"`
	Meta      		MenuMeta   		`json:"meta"`
	Children  		[]TreeList 		`json:"children"`
}

type UserRouteMenu struct {
	RouteTree 		[]TreeList		`json:"route_tree"`
	LeftMenuTree 	[]TreeList		`json:"left_menu_tree"`
}

type RouteMenu struct {
	Id 				uint64 			`json:"id"`
	ParentId 		uint64 			`json:"parent_id"`
	Type 			uint8       	`json:"type"`
	Level   		uint 			`json:"level"`
	Path 			string 			`json:"path"`
	Name 			string 			`json:"name"`
	Hidden 			bool 			`json:"hidden"`
	Component 		string 			`json:"component"`
	Icon 			string 			`json:"icon"`
	Code			string			`json:"code"`
	KeepAlive   	bool   			`json:"keep_alive"`
	Title       	string 			`json:"title"`
	IsLink      	bool 			`json:"is_link"`
}