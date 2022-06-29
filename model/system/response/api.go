package response

import "time"

type ApiGroupData struct {
	ID 				uint64     	`json:"id"`
	ApiGroup 		string      `json:"api_group"`
}

type ApiTree struct {
	ID 				uint64     	`json:"id"`
	Value 			uint64      `json:"value"`
	Title 			string      `json:"label"`
	Children   		[]ApiTree 	`json:"children"`
}

type ApiList struct {
	ID        		uint64 			`json:"id"`
	CreatedAt 		time.Time 		`json:"created_at"`
	ParentId 		uint64      	`json:"parent_id"`
	Description    	string			`json:"description"`
	Path 			string      	`json:"path"`
	Method 			string 			`json:"method"`
	Children   		[]ApiList 		`json:"children"`
}