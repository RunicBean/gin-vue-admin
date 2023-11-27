package layout_one

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CompanyRouter struct {
}

// InitCompanyRouter 初始化 公司 路由信息
func (s *CompanyRouter) InitCompanyRouter(Router *gin.RouterGroup) {
	companyRouter := Router.Group("company").Use(middleware.OperationRecord())
	companyRouterWithoutRecord := Router.Group("company")
	var companyApi = v1.ApiGroupApp.Layout_oneApiGroup.CompanyApi
	{
		companyRouter.POST("createCompany", companyApi.CreateCompany)             // 新建公司
		companyRouter.DELETE("deleteCompany", companyApi.DeleteCompany)           // 删除公司
		companyRouter.DELETE("deleteCompanyByIds", companyApi.DeleteCompanyByIds) // 批量删除公司
		companyRouter.PUT("updateCompany", companyApi.UpdateCompany)              // 更新公司
	}
	{
		companyRouterWithoutRecord.GET("findCompany", companyApi.FindCompany) // 根据ID获取公司
		companyRouterWithoutRecord.GET("findCName", companyApi.FindCompanyByPath)
		companyRouterWithoutRecord.GET("getCompanyList", companyApi.GetCompanyList) // 获取公司列表
	}
}
