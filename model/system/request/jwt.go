package request

import "github.com/dgrijalva/jwt-go"

type LoginClaims struct {
	ID          uint64 `json:"id"`
	RoleId 		uint64 `json:"role_id"`
	Username    string `json:"username"`
	BufferTime  int64  `json:"buffer_time"`
	jwt.StandardClaims
}