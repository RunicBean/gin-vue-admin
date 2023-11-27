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

type ItemDataApi struct {
}

var itemDtaService = service.ServiceGroupApp.Layout_oneServiceGroup.ItemDataService

// CreateItemData 创建项目数据
// @Tags ItemData
// @Summary 创建项目数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body layout_one.ItemData true "创建项目数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /itemDta/createItemData [post]
func (itemDtaApi *ItemDataApi) CreateItemData(c *gin.Context) {
	var itemDta layout_one.ItemData
	err := c.ShouldBindJSON(&itemDta)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"ImageLink": {utils.NotEmpty()},
		"Company":   {utils.NotEmpty()},
		"Title":     {utils.NotEmpty()},
	}
	if err := utils.Verify(itemDta, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := itemDtaService.CreateItemData(&itemDta); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteItemData 删除项目数据
// @Tags ItemData
// @Summary 删除项目数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body layout_one.ItemData true "删除项目数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /itemDta/deleteItemData [delete]
func (itemDtaApi *ItemDataApi) DeleteItemData(c *gin.Context) {
	var itemDta layout_one.ItemData
	err := c.ShouldBindJSON(&itemDta)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := itemDtaService.DeleteItemData(itemDta); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteItemDataByIds 批量删除项目数据
// @Tags ItemData
// @Summary 批量删除项目数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除项目数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /itemDta/deleteItemDataByIds [delete]
func (itemDtaApi *ItemDataApi) DeleteItemDataByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := itemDtaService.DeleteItemDataByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateItemData 更新项目数据
// @Tags ItemData
// @Summary 更新项目数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body layout_one.ItemData true "更新项目数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /itemDta/updateItemData [put]
func (itemDtaApi *ItemDataApi) UpdateItemData(c *gin.Context) {
	var itemDta layout_one.ItemData
	err := c.ShouldBindJSON(&itemDta)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"ImageLink": {utils.NotEmpty()},
		"Company":   {utils.NotEmpty()},
		"Title":     {utils.NotEmpty()},
	}
	if err := utils.Verify(itemDta, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := itemDtaService.UpdateItemData(itemDta); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindItemData 用id查询项目数据
// @Tags ItemData
// @Summary 用id查询项目数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query layout_one.ItemData true "用id查询项目数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /itemDta/findItemData [get]
func (itemDtaApi *ItemDataApi) FindItemData(c *gin.Context) {
	var itemDta layout_one.ItemData
	err := c.ShouldBindQuery(&itemDta)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reitemDta, err := itemDtaService.GetItemData(itemDta.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reitemDta": reitemDta}, c)
	}
}

// GetItemDataList 分页获取项目数据列表
// @Tags ItemData
// @Summary 分页获取项目数据列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query layout_oneReq.ItemDataSearch true "分页获取项目数据列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /itemDta/getItemDataList [get]
func (itemDtaApi *ItemDataApi) GetItemDataList(c *gin.Context) {
	var pageInfo layout_oneReq.ItemDataSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := itemDtaService.GetItemDataInfoList(pageInfo); err != nil {
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

func (itemDtaApi *ItemDataApi) GetItemDataListByCompany(c *gin.Context) {
	var pageInfo layout_oneReq.ItemDataSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := itemDtaService.GetItemDataInfoListByCompany(pageInfo); err != nil {
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
