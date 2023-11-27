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

type CarouselDataApi struct {
}

var caroDtaService = service.ServiceGroupApp.Layout_oneServiceGroup.CarouselDataService

// CreateCarouselData 创建轮播图数据
// @Tags CarouselData
// @Summary 创建轮播图数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body layout_one.CarouselData true "创建轮播图数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /caroDta/createCarouselData [post]
func (caroDtaApi *CarouselDataApi) CreateCarouselData(c *gin.Context) {
	var caroDta layout_one.CarouselData
	err := c.ShouldBindJSON(&caroDta)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"ImageLink": {utils.NotEmpty()},
		"Company":   {utils.NotEmpty()},
	}
	if err := utils.Verify(caroDta, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := caroDtaService.CreateCarouselData(&caroDta); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCarouselData 删除轮播图数据
// @Tags CarouselData
// @Summary 删除轮播图数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body layout_one.CarouselData true "删除轮播图数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /caroDta/deleteCarouselData [delete]
func (caroDtaApi *CarouselDataApi) DeleteCarouselData(c *gin.Context) {
	var caroDta layout_one.CarouselData
	err := c.ShouldBindJSON(&caroDta)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := caroDtaService.DeleteCarouselData(caroDta); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCarouselDataByIds 批量删除轮播图数据
// @Tags CarouselData
// @Summary 批量删除轮播图数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除轮播图数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /caroDta/deleteCarouselDataByIds [delete]
func (caroDtaApi *CarouselDataApi) DeleteCarouselDataByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := caroDtaService.DeleteCarouselDataByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCarouselData 更新轮播图数据
// @Tags CarouselData
// @Summary 更新轮播图数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body layout_one.CarouselData true "更新轮播图数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /caroDta/updateCarouselData [put]
func (caroDtaApi *CarouselDataApi) UpdateCarouselData(c *gin.Context) {
	var caroDta layout_one.CarouselData
	err := c.ShouldBindJSON(&caroDta)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"ImageLink": {utils.NotEmpty()},
		"Company":   {utils.NotEmpty()},
	}
	if err := utils.Verify(caroDta, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := caroDtaService.UpdateCarouselData(caroDta); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCarouselData 用id查询轮播图数据
// @Tags CarouselData
// @Summary 用id查询轮播图数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query layout_one.CarouselData true "用id查询轮播图数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /caroDta/findCarouselData [get]
func (caroDtaApi *CarouselDataApi) FindCarouselData(c *gin.Context) {
	var caroDta layout_one.CarouselData
	err := c.ShouldBindQuery(&caroDta)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recaroDta, err := caroDtaService.GetCarouselData(caroDta.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recaroDta": recaroDta}, c)
	}
}

// GetCarouselDataList 分页获取轮播图数据列表
// @Tags CarouselData
// @Summary 分页获取轮播图数据列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query layout_oneReq.CarouselDataSearch true "分页获取轮播图数据列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /caroDta/getCarouselDataList [get]
func (caroDtaApi *CarouselDataApi) GetCarouselDataList(c *gin.Context) {
	var pageInfo layout_oneReq.CarouselDataSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := caroDtaService.GetCarouselDataInfoList(pageInfo); err != nil {
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

func (caroDtaApi *CarouselDataApi) GetCarouselDataListByCompany(c *gin.Context) {
	var pageInfo layout_oneReq.CarouselDataSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := caroDtaService.GetCarouselDataInfoListByCompany(pageInfo); err != nil {
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
