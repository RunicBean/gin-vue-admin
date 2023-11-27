package layout_one

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/layout_one"
	layout_oneReq "github.com/flipped-aurora/gin-vue-admin/server/model/layout_one/request"
)

type CarouselDataService struct {
}

// CreateCarouselData 创建轮播图数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (caroDtaService *CarouselDataService) CreateCarouselData(caroDta *layout_one.CarouselData) (err error) {
	err = global.GVA_DB.Create(caroDta).Error
	return err
}

// DeleteCarouselData 删除轮播图数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (caroDtaService *CarouselDataService) DeleteCarouselData(caroDta layout_one.CarouselData) (err error) {
	err = global.GVA_DB.Delete(&caroDta).Error
	return err
}

// DeleteCarouselDataByIds 批量删除轮播图数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (caroDtaService *CarouselDataService) DeleteCarouselDataByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]layout_one.CarouselData{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCarouselData 更新轮播图数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (caroDtaService *CarouselDataService) UpdateCarouselData(caroDta layout_one.CarouselData) (err error) {
	err = global.GVA_DB.Save(&caroDta).Error
	return err
}

// GetCarouselData 根据id获取轮播图数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (caroDtaService *CarouselDataService) GetCarouselData(id uint) (caroDta layout_one.CarouselData, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&caroDta).Error
	return
}

// GetCarouselDataInfoList 分页获取轮播图数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (caroDtaService *CarouselDataService) GetCarouselDataInfoList(info layout_oneReq.CarouselDataSearch) (list []layout_one.CarouselData, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&layout_one.CarouselData{})
	var caroDtas []layout_one.CarouselData
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

	err = db.Find(&caroDtas).Error
	return caroDtas, total, err
}

func (caroDtaService *CarouselDataService) GetCarouselDataInfoListByCompany(info layout_oneReq.CarouselDataSearch) (list []layout_one.CarouselData, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&layout_one.CarouselData{})
	var caroDtas []layout_one.CarouselData
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("company = ? AND created_at BETWEEN ? AND ?", info.Company, info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&caroDtas).Error
	return caroDtas, total, err
}
