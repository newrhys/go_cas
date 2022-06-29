package response

type IconList struct {
	ID				uint64			`json:"id"`
	ParentId 		uint64 			`json:"parent_id"`
	Name 			string     		`json:"name"`
	Children   		[]IconList 		`json:"children"`
}
