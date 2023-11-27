package layout_one

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ItemDataRouter struct {
}

// InitItemDataRouter 初始化 项目数据 路由信息
func (s *ItemDataRouter) InitItemDataRouter(Router *gin.RouterGroup) {
	itemDtaRouter := Router.Group("itemDta").Use(middleware.OperationRecord())
	itemDtaRouterWithoutRecord := Router.Group("itemDta")
	var itemDtaApi = v1.ApiGroupApp.Layout_oneApiGroup.ItemDataApi
	{
		itemDtaRouter.POST("createItemData", itemDtaApi.CreateItemData)             // 新建项目数据
		itemDtaRouter.DELETE("deleteItemData", itemDtaApi.DeleteItemData)           // 删除项目数据
		itemDtaRouter.DELETE("deleteItemDataByIds", itemDtaApi.DeleteItemDataByIds) // 批量删除项目数据
		itemDtaRouter.PUT("updateItemData", itemDtaApi.UpdateItemData)              // 更新项目数据
	}
	{
		itemDtaRouterWithoutRecord.GET("findItemData", itemDtaApi.FindItemData)       // 根据ID获取项目数据
		itemDtaRouterWithoutRecord.GET("getItemDataList", itemDtaApi.GetItemDataList) // 获取项目数据列表
		itemDtaRouterWithoutRecord.GET("getItemDataListC", itemDtaApi.GetItemDataListByCompany)
	}
}
