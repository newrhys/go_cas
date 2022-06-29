package cms

import (
	"fmt"
	"wave-admin/global"
	"wave-admin/model/cms"
	"wave-admin/model/cms/request"
	"wave-admin/model/cms/response"
	"wave-admin/utils"
)

type TagService struct{}

func (tagService *TagService) AddTag(category request.ReqTag) (err error) {
	add := cms.CmsTag{
		Status: category.Status,
		Name: category.Name,
		Sort: category.Sort,
	}
	err = global.GnDb.Create(&add).Error

	return err
}

func (tagService *TagService) UpdateTag(tag request.ReqTag, id int) (err error) {
	update := cms.CmsTag{
		Status:    	tag.Status,
		Name:		tag.Name,
		Sort:      	tag.Sort,
	}
	toMap,err := utils.ToMap(&update, "json", "children")
	if err != nil {
		return err
	}
	err = global.GnDb.Model(&cms.CmsTag{}).Where("id=?", id).Updates(toMap).Error

	return err
}

func (tagService *TagService) DeleteTag(id int) (err error) {
	var count int64
	global.GnDb.Where("tag_id", id).Model(&cms.CmsArticleTag{}).Count(&count)
	if count > 0 {
		return fmt.Errorf("请删除此标签下的文章！")
	} else {
		err = global.GnDb.Delete(&cms.CmsTag{}, "id = ?", id).Error
		return err
	}
}

func (tagService *TagService) GetSort() (tag cms.CmsTag, err error)  {
	tag = cms.CmsTag{}
	err = global.GnDb.Order("sort DESC").Find(&tag).Error

	return tag,err
}

func (tagService *TagService) GetTagList(info request.TagList) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.CurrentPage - 1)
	var tags []cms.CmsTag
	db := global.GnDb.Model(&cms.CmsTag{}).Where("status = ?", info.Status).Order("sort")
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%").Find(&tags)
	}
	err = db.Count(&total).Error
	if err != nil {
		return err, tags, total
	}

	err = db.Limit(limit).Offset(offset).Find(&tags).Error

	return err, tags, total
}

func (tagService *TagService)SelectTagList() (err error, list []response.SelectTagList) {
	var tags []cms.CmsTag
	err = global.GnDb.Model(&cms.CmsTag{}).Where("status = ?", 1).Order("sort").Find(&tags).Error
	for _, tag := range tags {
		list = append(list, response.SelectTagList{
			ID:   tag.ID,
			Name: tag.Name,
		})
	}

	return err, list
}