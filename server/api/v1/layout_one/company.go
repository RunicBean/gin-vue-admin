package layout_one

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/layout_one"
	layout_oneReq "github.com/flipped-aurora/gin-vue-admin/server/model/layout_one/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CompanyApi struct {
}

var companyService = service.ServiceGroupApp.Layout_oneServiceGroup.CompanyService

// CreateCompany 创建公司
// @Tags Company
// @Summary 创建公司
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body layout_one.Company true "创建公司"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /company/createCompany [post]
func (companyApi *CompanyApi) CreateCompany(c *gin.Context) {
	var company layout_one.Company
	err := c.ShouldBindJSON(&company)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
		"Path": {utils.NotEmpty()},
	}
	if err := utils.Verify(company, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := companyService.CreateCompany(&company); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCompany 删除公司
// @Tags Company
// @Summary 删除公司
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body layout_one.Company true "删除公司"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /company/deleteCompany [delete]
func (companyApi *CompanyApi) DeleteCompany(c *gin.Context) {
	var company layout_one.Company
	err := c.ShouldBindJSON(&company)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := companyService.DeleteCompany(company); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCompanyByIds 批量删除公司
// @Tags Company
// @Summary 批量删除公司
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除公司"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /company/deleteCompanyByIds [delete]
func (companyApi *CompanyApi) DeleteCompanyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := companyService.DeleteCompanyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCompany 更新公司
// @Tags Company
// @Summary 更新公司
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body layout_one.Company true "更新公司"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /company/updateCompany [put]
func (companyApi *CompanyApi) UpdateCompany(c *gin.Context) {
	var company layout_one.Company
	err := c.ShouldBindJSON(&company)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
		"Path": {utils.NotEmpty()},
	}
	if err := utils.Verify(company, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := companyService.UpdateCompany(company); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCompany 用id查询公司
// @Tags Company
// @Summary 用id查询公司
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query layout_one.Company true "用id查询公司"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /company/findCompany [get]
func (companyApi *CompanyApi) FindCompany(c *gin.Context) {
	var company layout_one.Company
	err := c.ShouldBindQuery(&company)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recompany, err := companyService.GetCompany(company.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recompany": recompany}, c)
	}
}

func (companyApi *CompanyApi) FindCompanyByPath(c *gin.Context) {
	var company layout_one.Company
	err := c.ShouldBindQuery(&company)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recompany, err := companyService.GetCompanyByPath(company.Path); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recompany": recompany}, c)
	}
}

// GetCompanyList 分页获取公司列表
// @Tags Company
// @Summary 分页获取公司列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query layout_oneReq.CompanySearch true "分页获取公司列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /company/getCompanyList [get]
func (companyApi *CompanyApi) GetCompanyList(c *gin.Context) {
	var pageInfo layout_oneReq.CompanySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := companyService.GetCompanyInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
