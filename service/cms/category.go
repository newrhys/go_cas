package cms

import (
	"fmt"
	"wave-admin/global"
	"wave-admin/model/cms"
	"wave-admin/model/cms/request"
	CmsRep "wave-admin/model/cms/response"
	"wave-admin/utils"
)

type CategoryService struct{}

func (categoryService *CategoryService) AddCategory(category request.ReqCategory) (err error) {
	add := cms.CmsCategory{
		Status: category.Status,
		ParentId: category.ParentId,
		Name: category.Name,
		Sort: category.Sort,
	}
	err = global.GnDb.Create(&add).Error

	return err
}

func (categoryService *CategoryService) UpdateCategory(category request.ReqCategory, id int) (err error) {
	update := cms.CmsCategory{
		Status:    	category.Status,
		ParentId:  	category.ParentId,
		Name:		category.Name,
		Sort:      	category.Sort,
	}
	toMap,err := utils.ToMap(&update, "json", "children")
	if err != nil {
		return err
	}
	err = global.GnDb.Model(&cms.CmsCategory{}).Where("id=?", id).Updates(toMap).Error

	return err
}

func (categoryService *CategoryService) DeleteCategory(id int) (err error) {
	var count int64
	global.GnDb.Where("category_id", id).Model(&cms.CmsArticle{}).Count(&count)
	if count > 0 {
		return fmt.Errorf("请删除此分类下的文章！")
	} else {
		err = global.GnDb.Delete(&cms.CmsCategory{}, "id = ?", id).Error
		return err
	}
}

func (categoryService *CategoryService) getCategoryList(categoryList []cms.CmsCategory, pid uint64) (treeList []CmsRep.CategoryList) {
	for _, v := range categoryList {
		if v.ParentId == pid {
			child := categoryService.getCategoryList(categoryList, v.ID)
			node := CmsRep.CategoryList{
				ID:        v.ID,
				CreatedAt: v.CreatedAt,
				Status:    v.Status,
				ParentId:  v.ParentId,
				Name:      v.Name,
				Sort:      v.Sort,
			}
			node.Children = child

			treeList = append(treeList, node)
		}
	}
	return treeList
}

func (categoryService *CategoryService) GetCategoryInfoList(info request.CategoryList) (err error, list interface{}) {
	var categories []cms.CmsCategory
	db := global.GnDb.Model(&cms.CmsCategory{}).Where("status = ?", info.Status).Order("sort")
	if info.Name != "" {
		err = db.Where("name LIKE ?", "%"+info.Name+"%").Find(&categories).Error
		return err, categories
	} else {
		err = db.Find(&categories).Error
		if info.Status == "1" {
			tree := categoryService.getCategoryList(categories, 0)
			return err, tree
		} else {
			return err, categories
		}
	}
}

func (categoryService *CategoryService) getParentTree(cateList []cms.CmsCategory, pid uint64) (treeList []CmsRep.ParentCategory) {
	for _, v := range cateList {
		if v.ParentId == pid {
			child := categoryService.getParentTree(cateList, v.ID)
			node := CmsRep.ParentCategory{
				ID:       v.ID,
				Value:    v.ID,
				Title:    v.Name,
				Children: nil,
			}
			node.Children = child

			treeList = append(treeList, node)
		}
	}
	return treeList
}

func (categoryService *CategoryService) GetParentList() (list interface{}, err error) {
	db := global.GnDb.Model(&cms.CmsCategory{})
	var categoryList []cms.CmsCategory
	err = db.Order("sort").Where("status = ?", "1").Find(&categoryList).Error
	if err != nil {
		return nil,err
	}

	var rootTree []CmsRep.ParentCategory
	rootTree = append(rootTree, CmsRep.ParentCategory{
		ID:       0,
		Value:    0,
		Title:    "根分类",
		Children: nil,
	})

	tree := categoryService.getParentTree(categoryList, 0)

	rootTree = append(rootTree, tree...)

	return rootTree,err
}

func (categoryService *CategoryService) GetSort(id int) (category cms.CmsCategory, err error)  {
	category = cms.CmsCategory{}
	err = global.GnDb.Order("sort DESC").Where("parent_id = ?", id).Find(&category).Error

	return category,err
}