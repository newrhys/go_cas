package system

import (
	"errors"
	"fmt"
	"github.com/mojocn/base64Captcha"
	uuid "github.com/satori/go.uuid"
	"wave-admin/global"
	"wave-admin/model/system"
	"wave-admin/model/system/request"
	SystemRep "wave-admin/model/system/response"
	"wave-admin/utils"
)

type UserService struct{}

var store = base64Captcha.DefaultMemStore

var tempMenus = []SystemRep.TreeList{}

func (userService *UserService) Verify(id string, val string) bool {
	if id == "" || val == ""{
		return false
	}
	// 同时在内存清理掉这个图片
	return store.Verify(id, val, true)
}

// 添加用户
func (userService *UserService) AddUser(user request.AddUser) (err error) {
	username := user.Username
	password := user.Password
	u,_ := userService.UserByName(username)
	if u.ID != 0 {
		return fmt.Errorf("用户已存在！")
	}
	salt := utils.RandomString(5)
	// 创建用户
	newUser := system.SysUser{
		UUID: uuid.NewV4(),
		Username: username,
		Password: utils.MakePasswd(password, salt),
		Nickname: user.Nickname,
		Salt: salt,
		Avatar: "/uploads/default/logo.png",
		Status: user.Status,
		Mobile: user.Mobile,
		Email: user.Email,
	}
	err = global.GnDb.Create(&newUser).Error

	return err
}

// 登录
func (userService *UserService) Login(login request.Login) (user system.SysUser,err error) {
	// 获取参数
	username := login.Username
	password := login.Password
	user,err = userService.UserByName(username)
	if err == nil {
		ret := utils.ValidatePasswd(password, user.Salt, user.Password)
		if ret {
			return user, nil
		} else {
			return user, errors.New("密码错误！")
		}
	} else {
		return user,err
	}
}

// 根据用户名获得用户
func (userService *UserService) UserByName(username string) (user system.SysUser, err error) {
	user = system.SysUser{}
	err = global.GnDb.Where("username = ?", username).First(&user).Error
	user.Avatar = utils.TransformImageUrl(user.Avatar)

	return user,err
}

func (userService *UserService) GetUser(id uint64) (userInfo SystemRep.UserInfo, err error) {
	var	user system.SysUser
	err = global.GnDb.First(&user, id).Error
	if err != nil {
		return userInfo, err
	}

	var userRole system.SysUserRole
	err = global.GnDb.Model(&system.SysUserRole{}).Where("user_id = ?", user.ID).Find(&userRole).Error
	if err != nil {
		return userInfo, err
	}

	var role system.SysRole
	err = global.GnDb.Find(&role, userRole.RoleId).Error
	if err != nil {
		return userInfo, err
	}

	var results []SystemRep.RoleMenuJoinResult
	err = global.GnDb.Model(&system.SysRoleMenu{}).Select("sys_menus.code").Joins("left join sys_menus on sys_role_menus.menu_id = sys_menus.id").Where("sys_role_menus.role_id = ?", userRole.RoleId).Scan(&results).Error
	if err != nil {
		return userInfo, err
	}

	var permissions []string
	for _, roleMenu := range results {
		permissions = append(permissions, roleMenu.Code)
	}

	userInfo = SystemRep.UserInfo{
		Id:       		user.ID,
		UUID:     		user.UUID,
		Username: 		user.Username,
		Nickname: 		user.Nickname,
		Avatar:   		utils.TransformImageUrl(user.Avatar),
		Status:   		user.Status,
		Mobile:   		user.Mobile,
		Email:    		user.Email,
		Permissions: 	permissions,
		RoleName: 		role.RoleName,
		CreatedAt:      user.CreatedAt,
	}

	return userInfo,err
}

func (userService *UserService) GetUserList(info request.UserList) (err error, list interface{}, total int64) {
	if info.Status == "" {
		info.Status = "1"
	}

	limit := info.PageSize
	offset := info.PageSize * (info.CurrentPage - 1)
	db := global.GnDb.Model(&system.SysUser{}).Where("status = ?", info.Status)
	var userList []system.SysUser
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return err, userList, total
	}

	err = db.Limit(limit).Offset(offset).Find(&userList).Error

	var pageList []SystemRep.UserInfo
	for _, user := range userList {
		user.Avatar = utils.TransformImageUrl(user.Avatar)
		u := SystemRep.UserInfo{
			Id:       user.ID,
			UUID:     user.UUID,
			Username: user.Username,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
			Status:   user.Status,
			Mobile:   user.Mobile,
			Email:    user.Email,
		}
		pageList = append(pageList, u)
	}

	return err, pageList, total
}

func (userService *UserService) UpdateUser(reqUser request.UpdateUser, id int) (err error) {
	username := reqUser.Username
	password := reqUser.Password

	user := system.SysUser{}
	err = global.GnDb.First(&user, id).Error
	if user.Username != username {
		u,_ := userService.UserByName(username)
		if u.ID != 0 {
			return fmt.Errorf("用户已存在！")
		}
	}

	salt := utils.RandomString(5)
	update := system.SysUser{
		Username: reqUser.Username,
		Nickname: reqUser.Nickname,
		Status:   reqUser.Status,
		Mobile:   reqUser.Mobile,
		Email:    reqUser.Email,
	}
	if len(password) > 0 {
		update.Salt = salt
		update.Password = utils.MakePasswd(password, salt)
	}
	err = global.GnDb.Model(&system.SysUser{}).Where("id = ?", id).Updates(update).Error
	return err
}

func (userService *UserService) DeleteUser(id int) (err error) {
	var user system.SysUser
	err = global.GnDb.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	err = global.GnDb.Delete(&system.SysUserRole{}, "user_id = ?", id).Error

	return err
}

func (userService *UserService) GetRoleIdByUserId(id int) (user system.SysUserRole, err error) {
	var userRole system.SysUserRole
	err = global.GnDb.Where("user_id = ?", id).Find(&userRole).Error
	return userRole, err
}

func (userService *UserService) AssignRole(assignRole request.AssignRole) (err error) {
	var userRole system.SysUserRole

	err = global.GnDb.Where("user_id = ?", assignRole.UserId).Find(&userRole).Error

	if err != nil {
		return err
	}

	db := global.GnDb.Model(&system.SysUserRole{})
	if userRole.UserId > 0 {
		err = db.Where("user_id = ?", assignRole.UserId).Updates(&system.SysUserRole{
			RoleId:  assignRole.RoleId,
		}).Error
	} else {
		err = db.Create(&system.SysUserRole{
			UserId:  assignRole.UserId,
			RoleId:  assignRole.RoleId,
		}).Error
	}

	if err != nil {
		return err
	}

	return nil
}

func (userService *UserService) getRouteMenuTree(menuList []SystemRep.RouteMenu, pid uint64, isRoute bool)(treeList []SystemRep.TreeList) {
	for _, v := range menuList {
		if v.ParentId == pid {
			node := SystemRep.TreeList{
				Id: 		v.Id,
				ParentId:  	v.ParentId,
				Type:       v.Type,
				Component:  v.Component,
				Name: 		v.Name,
				Path:       v.Path,
				Hidden:     v.Hidden,
				IsLink: 	v.IsLink,
				Meta: SystemRep.MenuMeta{
					Icon:       v.Icon,
					Title:      v.Title,
					Code:       v.Code,
					KeepAlive:  v.KeepAlive,
				},
			}
			if isRoute && v.Type == 1 && v.Level == 2 {
				tempMenus = append(tempMenus, node)
			} else {
				child := userService.getRouteMenuTree(menuList, v.Id, isRoute)
				node.Children = child
				treeList = append(treeList, node)
			}
		}
	}
	return treeList
}

func (userService *UserService) GetRouteMenuList(id uint64) (list interface{}, err error) {
	var userRole system.SysUserRole
	err = global.GnDb.Model(&system.SysUserRole{}).Where("user_id = ?", id).Find(&userRole).Error
	if err != nil {
		return nil, err
	}

	var results []SystemRep.RouteMenu
	err = global.GnDb.Model(&system.SysRoleMenu{}).Select("sys_menus.*").Joins("left join sys_menus on sys_role_menus.menu_id = sys_menus.id").Where("sys_role_menus.role_id = ?", userRole.RoleId).Where("sys_menus.type < ?", 2).Order("sys_menus.sort").Scan(&results).Error
	if err != nil {
		return nil, err
	}

	routeTree := userService.getRouteMenuTree(results, 0, true)
	//log.Println(tempMenus)
	for _, temp := range tempMenus {
		for i, v := range routeTree {
			if v.Children != nil {
				for _, val := range v.Children {
					if val.Id == temp.ParentId {
						routeTree[i].Children = append(v.Children, temp)
					}
				}
			}
		}
	}

	leftMenuTree := userService.getRouteMenuTree(results, 0, false)

	return SystemRep.UserRouteMenu{
		RouteTree:    routeTree,
		LeftMenuTree: leftMenuTree,
	}, err
}

func (userService *UserService) ChangePassword(pwd request.ChangePassword, id uint64) (err error) {
	var	user system.SysUser
	err = global.GnDb.First(&user, id).Error
	if err != nil {
		return err
	}

	ret := utils.ValidatePasswd(pwd.OldPassword, user.Salt, user.Password)
	if !ret {
		return errors.New("旧密码错误！")
	}

	if pwd.Password != pwd.ConfirmPassword {
		return errors.New("两次密码输入错误！")
	}

	update := system.SysUser{
		Password: utils.MakePasswd(pwd.Password, user.Salt),
	}

	err = global.GnDb.Model(&system.SysUser{}).Where("id = ?", id).Updates(update).Error
	return err
}