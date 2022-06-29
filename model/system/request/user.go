package request

import (
	"wave-admin/model/common/request"
)

type Login struct {
	Username   			string 		`json:"username" binding:"required,max=16,min=3" label:"用户名"`
	Password 			string 		`json:"password" binding:"required,max=16,min=6" label:"密码"`
	//CaptchaId	string		`json:"captcha_id" binding:"required"`
	//VerifyCode	string		`json:"verify_code" binding:"required"`
}

type Register struct {
	Login
}

type UserList struct {
	Username			string		`form:"username"`
	Status 				string		`form:"status"`
	request.PageInfo
}

type AddUser struct {
	Username   			string 		`json:"username" binding:"required,max=16,min=2" label:"用户名"`
	Password 			string 		`json:"password" binding:"required,max=16,min=6" label:"密码"`
	Nickname 			string 		`json:"nickname" binding:"required,max=16,min=2" label:"昵称"`
	Status 				uint8		`json:"status" binding:"required" label:"状态"`
	Mobile 				string		`json:"mobile"`
	Email 				string		`json:"email"`
}

type UpdateUser struct {
	Username   			string 		`json:"username" binding:"required,max=16,min=2" label:"用户名"`
	Password 			string 		`json:"password"`
	Nickname 			string 		`json:"nickname" binding:"required,max=16,min=2" label:"昵称"`
	Status 				uint8		`json:"status"`
	Mobile 				string		`json:"mobile"`
	Email 				string		`json:"email"`
}

type AssignRole struct {
	UserId 				uint64 		`json:"user_id"`
	RoleId 				uint64 		`json:"role_id"`
}

type ChangePassword struct {
	OldPassword 		string 		`json:"old_password" binding:"required,max=16,min=6" label:"旧密码"`
	Password 			string 		`json:"password" binding:"required,max=16,min=6" label:"新密码"`
	ConfirmPassword		string 		`json:"confirm_password" binding:"required,max=16,min=6" label:"确认密码"`
}