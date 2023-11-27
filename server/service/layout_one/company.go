package layout_one

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/layout_one"
	layout_oneReq "github.com/flipped-aurora/gin-vue-admin/server/model/layout_one/request"
)

type CompanyService struct {
}

// CreateCompany 创建公司记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyService *CompanyService) CreateCompany(company *layout_one.Company) (err error) {
	err = global.GVA_DB.Create(company).Error
	return err
}

// DeleteCompany 删除公司记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyService *CompanyService) DeleteCompany(company layout_one.Company) (err error) {
	err = global.GVA_DB.Delete(&company).Error
	return err
}

// DeleteCompanyByIds 批量删除公司记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyService *CompanyService) DeleteCompanyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]layout_one.Company{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCompany 更新公司记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyService *CompanyService) UpdateCompany(company layout_one.Company) (err error) {
	err = global.GVA_DB.Save(&company).Error
	return err
}

// GetCompany 根据id获取公司记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyService *CompanyService) GetCompany(id uint) (company layout_one.Company, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&company).Error
	return
}

// GetCompanyByPath 根据id获取公司记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyService *CompanyService) GetCompanyByPath(path string) (company layout_one.Company, err error) {
	err = global.GVA_DB.Where("path = ?", path).First(&company).Error
	return
}

// GetCompanyInfoList 分页获取公司记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyService *CompanyService) GetCompanyInfoList(info layout_oneReq.CompanySearch) (list []layout_one.Company, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&layout_one.Company{})
	var companys []layout_one.Company
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

	err = db.Find(&companys).Error
	return companys, total, err
}
