package response

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type LoginResponse struct {
	Token     	string        `json:"token"`
	ExpiresAt 	int64         `json:"expiresAt"`
}

type RoleMenuJoinResult struct {
	Code  string
}

type UserInfo struct {
	Id 				uint64 		`json:"id"`
	UUID 			uuid.UUID 	`json:"uuid"`
	Username 		string 		`json:"username"`
	Nickname 		string 		`json:"nickname"`
	Avatar 			string 		`json:"avatar"`
	Status 			uint8 		`json:"status"`
	Mobile 			string 		`json:"mobile"`
	Email 			string 		`json:"email"`
	Permissions     []string	`json:"permissions"`
	RoleName        string 		`json:"role_name"`
	CreatedAt 		time.Time 	`json:"created_at"`
}