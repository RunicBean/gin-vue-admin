package layout_one

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/layout_one"
	layout_oneReq "github.com/flipped-aurora/gin-vue-admin/server/model/layout_one/request"
)

type ItemDataService struct {
}

// CreateItemData 创建项目数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemDtaService *ItemDataService) CreateItemData(itemDta *layout_one.ItemData) (err error) {
	err = global.GVA_DB.Create(itemDta).Error
	return err
}

// DeleteItemData 删除项目数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemDtaService *ItemDataService) DeleteItemData(itemDta layout_one.ItemData) (err error) {
	err = global.GVA_DB.Delete(&itemDta).Error
	return err
}

// DeleteItemDataByIds 批量删除项目数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemDtaService *ItemDataService) DeleteItemDataByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]layout_one.ItemData{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateItemData 更新项目数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemDtaService *ItemDataService) UpdateItemData(itemDta layout_one.ItemData) (err error) {
	err = global.GVA_DB.Save(&itemDta).Error
	return err
}

// GetItemData 根据id获取项目数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemDtaService *ItemDataService) GetItemData(id uint) (itemDta layout_one.ItemData, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&itemDta).Error
	return
}

// GetItemDataInfoList 分页获取项目数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemDtaService *ItemDataService) GetItemDataInfoList(info layout_oneReq.ItemDataSearch) (list []layout_one.ItemData, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&layout_one.ItemData{})
	var itemDtas []layout_one.ItemData
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&itemDtas).Error
	return itemDtas, total, err
}

// GetItemDataInfoListByCompany 分页获取项目数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemDtaService *ItemDataService) GetItemDataInfoListByCompany(info layout_oneReq.ItemDataSearch) (list []layout_one.ItemData, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&layout_one.ItemData{})
	var itemDtas []layout_one.ItemData
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.ItemData.Company, info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Company != "" {
		db = db.Where("company = ?", info.Company)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&itemDtas).Error
	return itemDtas, total, err
}
