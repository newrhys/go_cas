package system

import (
	"fmt"
	"wave-admin/global"
	"wave-admin/model/system"
	"wave-admin/model/system/request"
	SystemRep "wave-admin/model/system/response"
)

type ApiService struct{}

func (apiService *ApiService) GetApi(id int) (api system.SysApi, err error) {
	api = system.SysApi{}
	err = global.GnDb.First(&api, id).Error

	return api,err
}

func (apiService *ApiService) GetApiByPath(path string, method string) (api system.SysApi, err error) {
	api = system.SysApi{}
	err = global.GnDb.Model(&system.SysApi{}).Where("path = ?", path).Where("method = ?", method).First(&api).Error

	return api,err
}

func (apiService *ApiService) AddApi(req request.ReqApi) (err error) {
	api := system.SysApi{
		ParentId:    req.ParentId,
		Description: req.Description,
		Path:        req.Path,
		Method:      req.Method,
	}
	err = global.GnDb.Create(&api).Error

	return err
}

func (apiService *ApiService) UpdateApi(req request.ReqApi, id int) (err error) {
	update := system.SysApi{
		ParentId:    req.ParentId,
		Description: req.Description,
		Path:        req.Path,
		Method:      req.Method,
	}
	err = global.GnDb.Model(&system.SysApi{}).Where("id=?", id).Updates(update).Error

	return err
}

func (apiService *ApiService) DeleteApi(id int) (err error) {
	var count int64
	global.GnDb.Where("api_id", id).Model(&system.SysRoleApi{}).Count(&count)
	if count > 0 {
		return fmt.Errorf("请删除角色API权限！")
	} else  {
		err = global.GnDb.Delete(&system.SysApi{}, "id = ?", id).Error
		return err
	}
}

func (apiService *ApiService) getApiList(apiList []system.SysApi, pid uint64) (treeList []SystemRep.ApiList) {
	for _, v := range apiList {
		if v.ParentId == pid {
			child := apiService.getApiList(apiList, v.ID)
			node := SystemRep.ApiList{
				ID:          v.ID,
				CreatedAt:   v.CreatedAt,
				ParentId:    v.ParentId,
				Description: v.Description,
				Path:        v.Path,
				Method:      v.Method,
				Children:    nil,
			}
			node.Children = child

			treeList = append(treeList, node)
		}
	}
	return treeList
}

func (apiService *ApiService) GetApiInfoList(info request.ApiList) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.CurrentPage - 1)
	db := global.GnDb.Model(&system.SysApi{})
	var apis []system.SysApi
	if info.Path != "" || info.Description != "" || info.Method != "" {
		if info.Path != "" {
			db = db.Where("path LIKE ?", "%"+info.Path+"%")
		}

		if info.Description != "" {
			db = db.Where("description LIKE ?", "%"+info.Description+"%")
		}

		if info.Method != "" {
			db = db.Where("method = ?", info.Method)
		}
		err = db.Count(&total).Error
		if err != nil {
			return err, apis, total
		}
		err = db.Offset(offset).Limit(limit).Find(&apis).Error
		return err, apis, total
	} else {
		total = 0
		err = db.Find(&apis).Error
		tree := apiService.getApiList(apis, 0)

		return err, tree, total
	}
}

func (apiService *ApiService) getParentTree(apiList []system.SysApi, pid uint64) (treeList []SystemRep.ApiTree) {
	for _, v := range apiList {
		if v.ParentId == pid {
			child := apiService.getParentTree(apiList, v.ID)
			node := SystemRep.ApiTree{
				ID:       v.ID,
				Value:    v.ID,
				Title:    v.Description,
				Children: nil,
			}
			node.Children = child

			treeList = append(treeList, node)
		}
	}
	return treeList
}

func (apiService *ApiService) GetTreeList() (list interface{}, err error) {
	var apis []system.SysApi
	err = global.GnDb.Model(&system.SysApi{}).Find(&apis).Error
	if err != nil {
		return nil, err
	}

	var rootTree []SystemRep.ApiTree
	rootTree = append(rootTree, SystemRep.ApiTree{
		ID:       0,
		Value:    0,
		Title:    "根API",
		Children: nil,
	})

	tree := apiService.getParentTree(apis, 0)

	rootTree = append(rootTree, tree...)

	return rootTree,err
}
