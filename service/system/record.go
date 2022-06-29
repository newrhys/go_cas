package system

import (
	"wave-admin/global"
	"wave-admin/model/system"
	"wave-admin/model/system/request"
)

type RecordService struct{}

func (recordService *RecordService) AddRecord(record system.SysRecord) (err error) {
	err = global.GnDb.Create(&record).Error
	return err
}

func (recordService *RecordService) DeleteRecord(id int) (err error) {
	err = global.GnDb.Delete(&system.SysRecord{}, "id = ?", id).Error
	return err
}

func (recordService *RecordService) GetRecordInfoList(info request.RecordList) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.CurrentPage - 1)
	// 创建db
	db := global.GnDb.Model(&system.SysRecord{})
	var records []system.SysRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Path != "" {
		db = db.Where("path LIKE ?", "%"+info.Path+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&records).Error
	return err, records, total
}
