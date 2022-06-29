package response

import "time"

type CategoryList struct {
	ID        		uint64 				`json:"id"`
	CreatedAt 		time.Time 			`json:"created_at"`
	Status 			uint8 				`json:"status"`
	ParentId 		uint64 				`json:"parent_id"`
	Name 			string     			`json:"name"`
	Sort 			int  				`json:"sort"`
	Children   		[]CategoryList		`json:"children"`
}

type ParentCategory struct {
	ID 				uint64     			`json:"id"`
	Value 			uint64      		`json:"value"`
	Title 			string      		`json:"label"`
	Children   		[]ParentCategory 	`json:"children"`
}