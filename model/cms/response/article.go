package response

type Article struct {
	ID        		uint64 		`json:"id"`
	Status 			string     	`json:"status"`
	Title       	string      `json:"title"`
	Img 			string 		`json:"img"`
	KeyWord			string 		`json:"key_word"`
	CategoryId 		uint64 		`json:"category_id"`
	Description    	string		`json:"description"`
	Content    		string		`json:"content"`
	Tag 			[]int		`json:"tag"`
}
