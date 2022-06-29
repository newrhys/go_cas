package system

import (
	"wave-admin/global"
	"wave-admin/model/system"
	"wave-admin/model/system/request"
	SystemRep "wave-admin/model/system/response"
)

type IconService struct{}

func (iconService *IconService) getIconList(iconList []system.SysIcon, pid uint64) (treeList []SystemRep.IconList) {
	for _, v := range iconList {
		if v.ParentId == pid {
			child := iconService.getIconList(iconList, v.ID)
			node := SystemRep.IconList{
				ID: 	   v.ID,
				ParentId:  v.ParentId,
				Name:      v.Name,
			}
			node.Children = child

			treeList = append(treeList, node)
		}
	}
	return treeList
}

func (iconService *IconService) inArray(need uint64, needArr []uint64) bool {
	for _,v := range needArr{
		if need == v {
			return true
		}
	}
	return false
}

func (iconService *IconService) GetIconInfoList(info request.IconList) (err error, list interface{}) {
	var icons []system.SysIcon
	db := global.GnDb.Model(&system.SysIcon{})
	if info.Name != "" {
		err = db.Where("name LIKE ?", "%"+info.Name+"%").Where("parent_id > ?", 0).Find(&icons).Error
		if err != nil {
			return err, nil
		}
		var ids []uint64
		for _, icon := range icons {
			exists := iconService.inArray(icon.ParentId, ids)
			if exists == false {
				ids = append(ids, icon.ParentId)
			}
		}
		//db.Find(&users, []int{1,2,3})
		// SELECT * FROM users WHERE id IN (1,2,3);
		var parentIcons []system.SysIcon
		err = global.GnDb.Model(&system.SysIcon{}).Find(&parentIcons, ids).Error
		if err != nil {
			return err, nil
		}
		parentIcons = append(parentIcons, icons...)

		tree := iconService.getIconList(parentIcons, 0)

		return err, tree
	} else {
		err = db.Find(&icons).Error
		tree := iconService.getIconList(icons, 0)
		return err, tree
	}
}