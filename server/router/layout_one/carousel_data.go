package layout_one

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CarouselDataRouter struct {
}

// InitCarouselDataRouter 初始化 轮播图数据 路由信息
func (s *CarouselDataRouter) InitCarouselDataRouter(Router *gin.RouterGroup) {
	caroDtaRouter := Router.Group("caroDta").Use(middleware.OperationRecord())
	caroDtaRouterWithoutRecord := Router.Group("caroDta")
	var caroDtaApi = v1.ApiGroupApp.Layout_oneApiGroup.CarouselDataApi
	{
		caroDtaRouter.POST("createCarouselData", caroDtaApi.CreateCarouselData)             // 新建轮播图数据
		caroDtaRouter.DELETE("deleteCarouselData", caroDtaApi.DeleteCarouselData)           // 删除轮播图数据
		caroDtaRouter.DELETE("deleteCarouselDataByIds", caroDtaApi.DeleteCarouselDataByIds) // 批量删除轮播图数据
		caroDtaRouter.PUT("updateCarouselData", caroDtaApi.UpdateCarouselData)              // 更新轮播图数据
	}
	{
		caroDtaRouterWithoutRecord.GET("findCarouselData", caroDtaApi.FindCarouselData)       // 根据ID获取轮播图数据
		caroDtaRouterWithoutRecord.GET("getCarouselDataList", caroDtaApi.GetCarouselDataList) // 获取轮播图数据列表
		caroDtaRouterWithoutRecord.GET("getCarouselDataListC", caroDtaApi.GetCarouselDataListByCompany)
	}
}
