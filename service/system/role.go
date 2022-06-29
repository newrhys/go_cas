package system

import (
	adapter "github.com/casbin/gorm-adapter/v3"
	"wave-admin/global"
	"wave-admin/model/system"
	"wave-admin/model/system/request"
)

type RoleService struct{}

func (roleService *RoleService) GetRole(id int) (role system.SysRole, err error) {
	role = system.SysRole{}
	err = global.GnDb.First(&role, id).Error

	return role,err
}

func (roleService *RoleService) AddRole(req request.ReqRole) (err error, role system.SysRole) {
	role = system.SysRole{
		RoleName: 		req.RoleName,
		Remark:        	req.Remark,
		Status:    		req.Status,
	}
	err = global.GnDb.Create(&role).Error

	return err,role
}

func (roleService *RoleService) UpdateRole(role request.ReqRole) (err error) {
	update := system.SysRole{
		RoleName: 		role.RoleName,
		Remark:        	role.Remark,
		Status:    		role.Status,
	}
	err = global.GnDb.Model(&system.SysRole{}).Where("id=?", role.ID).Updates(update).Error

	return err
}

func (roleService *RoleService) DeleteRole(id int) (err error) {
	err = global.GnDb.Where("role_id = ?", id).Unscoped().Delete(&system.SysRoleMenu{}).Error
	if err != nil {
		return err
	}
	err = global.GnDb.Where("role_id = ?", id).Unscoped().Delete(&system.SysRoleApi{}).Error
	if err != nil {
		return err
	}
	err = global.GnDb.Delete(&system.SysRole{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	return err
}

func (roleService *RoleService) GetRoleInfoList(info request.RoleList) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.CurrentPage - 1)
	db := global.GnDb.Model(&system.SysRole{}).Where("status = ?", info.Status)
	var roles []system.SysRole
	if info.RoleName != "" {
		db = db.Where("role_name LIKE ?", "%"+info.RoleName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return err, roles, total
	}

	err = db.Limit(limit).Offset(offset).Find(&roles).Error

	return err, roles, total
}

func (roleService *RoleService) GetAssignMenu(roleId int) (err error, list []uint64) {
	var roleMenus []system.SysRoleMenu
	err = global.GnDb.Where("role_id = ?", roleId).Find(&roleMenus).Error
	for _, rm := range roleMenus {
		list = append(list, rm.MenuId)
	}
	return err, list
}

func (roleService *RoleService) AssignSave(assignSave request.AssignSave) (err error, casbinInfos []request.CasbinInfo) {
	err = global.GnDb.Where("role_id = ?", assignSave.ID).Unscoped().Delete(&system.SysRoleMenu{}).Error
	if err != nil {
		return err, nil
	}

	err = global.GnDb.Where("role_id = ?", assignSave.ID).Unscoped().Delete(&system.SysRoleApi{}).Error
	if err != nil {
		return err, nil
	}

	err = global.GnDb.Where("v0 = ?", assignSave.ID).Unscoped().Delete(&adapter.CasbinRule{}).Error
	if err != nil {
		return err, nil
	}

	var roleMenus []system.SysRoleMenu
	db := global.GnDb.Model(&system.SysRoleMenu{})
	for _, menuId := range assignSave.MenuList {
		roleMenus = append(roleMenus, system.SysRoleMenu{
			RoleId:  assignSave.ID,
			MenuId:  uint64(menuId),
		})
	}
	err = db.Create(&roleMenus).Error
	if err != nil {
		return err, nil
	}

	var apis []system.SysApi
	err = global.GnDb.Model(&system.SysApi{}).Find(&apis, assignSave.ApiList).Error
	if err != nil {
		return err, nil
	}
	//log.Println(apis)

	var roleApis []system.SysRoleApi
	for _, api := range apis {
		roleApis = append(roleApis, system.SysRoleApi{
			RoleId:  assignSave.ID,
			ApiId:   api.ID,
		})
		//log.Println("api.Path=", api.Path)
		if api.Path != "" {
			casbinInfos = append(casbinInfos, request.CasbinInfo{
				Path:   api.Path,
				Method: api.Method,
			})
		}
	}
	db2 := global.GnDb.Model(&system.SysRoleApi{})
	err = db2.Create(&roleApis).Error
	if err != nil {
		return err, nil
	}

	return err, casbinInfos
}

func (roleService *RoleService) GetAssignApi(roleId int) (err error, list []uint64) {
	var roleApis []system.SysRoleApi
	err = global.GnDb.Where("role_id = ?", roleId).Find(&roleApis).Error
	for _, rm := range roleApis {
		list = append(list, rm.ApiId)
	}
	return err, list
}
