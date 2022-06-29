package system

import "wave-admin/service"

type ApiGroup struct {
	AttachApi
	MenuApi
	RoleApi
	UserApi
	IconApi
	SysApiApi
	RecordApi
}

var (
	attachService = service.ServiceGroupApp.SystemServiceGroup.AttachService
	jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
	menuService = service.ServiceGroupApp.SystemServiceGroup.MenuService
	roleService = service.ServiceGroupApp.SystemServiceGroup.RoleService
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
	iconService = service.ServiceGroupApp.SystemServiceGroup.IconService
	apiService = service.ServiceGroupApp.SystemServiceGroup.ApiService
	recordService = service.ServiceGroupApp.SystemServiceGroup.RecordService
	casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService
)
