package system

import (
	"fmt"
	"log"
	"wave-admin/global"
	"wave-admin/model/system"
	"wave-admin/model/system/request"
	SystemRep "wave-admin/model/system/response"
	"wave-admin/utils"
)

type MenuService struct{}

func (menuService *MenuService) GetMenu(id int) (menu system.SysMenu, err error) {
	menu = system.SysMenu{}
	err = global.GnDb.First(&menu, id).Error

	return menu,err
}

func (menuService *MenuService) fixMenuLevel(parentId uint64, level *uint) {
	if parentId > 0 {
		*level += 1
		menu := system.SysMenu{}
		global.GnDb.First(&menu, parentId)
		//log.Println(menu)
		if menu.ParentId > 0 {
			menuService.fixMenuLevel(menu.ParentId, level)
		}
	}
}

func (menuService *MenuService) AddMenu(menu request.ReqMenu) (err error) {
	var level uint = 0
	menuService.fixMenuLevel(menu.ParentId, &level)
	log.Println("level=", level)
	add := system.SysMenu{
		Status:    	menu.Status,
		ParentId:  	menu.ParentId,
		Type: 	   	menu.Type,
		Level:  	level,
		Path:      	menu.Path,
		Name:		menu.Name,
		Hidden:    	menu.Hidden != 0,
		Component: 	menu.Component,
		Sort:      	menu.Sort,
		Icon:      	menu.Icon,
		Code:      	menu.Code,
		KeepAlive: 	menu.KeepAlive != 0,
		Title:     	menu.Title,
		IsLink:    	menu.IsLink != 0,
	}
	err = global.GnDb.Create(&add).Error

	return err
}

func (menuService *MenuService) UpdateMenu(menu request.ReqMenu, id int) (err error) {
	var level uint = 0
	menuService.fixMenuLevel(menu.ParentId, &level)
	log.Println("level=", level)
	update := system.SysMenu{
		Status:    	menu.Status,
		ParentId:  	menu.ParentId,
		Type: 		menu.Type,
		Level:  	level,
		Path:      	menu.Path,
		Name:		menu.Name,
		Hidden:    	menu.Hidden != 0,
		Component: 	menu.Component,
		Sort:      	menu.Sort,
		Icon:      	menu.Icon,
		Code:      	menu.Code,
		KeepAlive: 	menu.KeepAlive != 0,
		Title:     	menu.Title,
		IsLink:    	menu.IsLink != 0,
	}
	toMap,err := utils.ToMap(&update, "json", "children")
	if err != nil {
		return err
	}
	err = global.GnDb.Model(&system.SysMenu{}).Where("id=?", id).Updates(toMap).Error

	return err
}

func (menuService *MenuService) DeleteMenu(id int) (err error) {
	var count int64
	global.GnDb.Where("menu_id", id).Model(&system.SysRoleMenu{}).Count(&count)
	if count > 0 {
		return fmt.Errorf("请删除角色菜单权限！")
	} else {
		err = global.GnDb.Delete(&system.SysMenu{}, "id = ?", id).Error
		return err
	}
}

func (menuService *MenuService) getMenuList(menuList []system.SysMenu, pid uint64) (treeList []SystemRep.MenuList) {
	for _, v := range menuList {
		if v.ParentId == pid {
			child := menuService.getMenuList(menuList, v.ID)
			node := SystemRep.MenuList{
				ID:        v.ID,
				CreatedAt: v.CreatedAt,
				Status:    v.Status,
				ParentId:  v.ParentId,
				Type:      v.Type,
				Path:      v.Path,
				Name:      v.Name,
				Hidden:    v.Hidden,
				Component: v.Component,
				Sort:      v.Sort,
				Icon:      v.Icon,
				Code:      v.Code,
				KeepAlive: v.KeepAlive,
				Title:     v.Title,
				IsLink:    v.IsLink,
				Children:  nil,
			}
			node.Children = child

			treeList = append(treeList, node)
		}
	}
	return treeList
}

func (menuService *MenuService) GetMenuInfoList(info request.MenuList) (err error, list interface{}) {
	var menus []system.SysMenu
	db := global.GnDb.Model(&system.SysMenu{}).Where("status = ?", info.Status).Order("sort")
	if info.Title != "" {
		err = db.Where("title LIKE ?", "%"+info.Title+"%").Find(&menus).Error
		return err, menus
	} else {
		err = db.Find(&menus).Error
		if info.Status == "1" {
			tree := menuService.getMenuList(menus, 0)
			return err, tree
		} else {
			return err, menus
		}
	}
}

func (menuService *MenuService) getParentTree(menuList []system.SysMenu, pid uint64) (treeList []SystemRep.ParentMenu) {
	for _, v := range menuList {
		if v.ParentId == pid {
			child := menuService.getParentTree(menuList, v.ID)
			node := SystemRep.ParentMenu{
				ID:       v.ID,
				Value:    v.ID,
				Title:    v.Title,
				Children: nil,
			}
			node.Children = child

			treeList = append(treeList, node)
		}
	}
	return treeList
}

func (menuService *MenuService) GetParentList() (list interface{}, err error) {
	db := global.GnDb.Model(&system.SysMenu{})
	var menu []system.SysMenu
	err = db.Order("sort").Where("status = ?", "1").Find(&menu).Error
	if err != nil {
		return nil, err
	}

	var rootTree []SystemRep.ParentMenu
	rootTree = append(rootTree, SystemRep.ParentMenu{
		ID:       0,
		Value:    0,
		Title:    "根目录",
		Children: nil,
	})

	tree := menuService.getParentTree(menu, 0)

	rootTree = append(rootTree, tree...)

	return rootTree,err
}

func (menuService *MenuService) fixChildrenMenu(pm *SystemRep.ParentMenu, menu system.SysMenu)  {
	if len(menu.Children) > 0 {
		for i, m := range menu.Children {
			tmp := SystemRep.ParentMenu{
				ID:       	m.ID,
				Value:  	m.ID,
				Title:    	m.Title,
				Children: 	nil,
			}
			pm.Children = append(pm.Children, tmp)
			menuService.fixChildrenMenu(&pm.Children[i], m)
		}
	}
}

func (menuService *MenuService) GetSort(id int) (menu system.SysMenu, err error) {
	menu = system.SysMenu{}
	err = global.GnDb.Order("sort DESC").Where("parent_id = ?", id).Find(&menu).Error

	return menu,err
}