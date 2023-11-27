import service from '@/utils/request'

// @Tags ItemData
// @Summary 创建项目数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ItemData true "创建项目数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /itemDta/createItemData [post]
export const createItemData = (data) => {
  return service({
    url: '/itemDta/createItemData',
    method: 'post',
    data
  })
}

// @Tags ItemData
// @Summary 删除项目数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ItemData true "删除项目数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /itemDta/deleteItemData [delete]
export const deleteItemData = (data) => {
  return service({
    url: '/itemDta/deleteItemData',
    method: 'delete',
    data
  })
}

// @Tags ItemData
// @Summary 批量删除项目数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除项目数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /itemDta/deleteItemData [delete]
export const deleteItemDataByIds = (data) => {
  return service({
    url: '/itemDta/deleteItemDataByIds',
    method: 'delete',
    data
  })
}

// @Tags ItemData
// @Summary 更新项目数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ItemData true "更新项目数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /itemDta/updateItemData [put]
export const updateItemData = (data) => {
  return service({
    url: '/itemDta/updateItemData',
    method: 'put',
    data
  })
}

// @Tags ItemData
// @Summary 用id查询项目数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ItemData true "用id查询项目数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /itemDta/findItemData [get]
export const findItemData = (params) => {
  return service({
    url: '/itemDta/findItemData',
    method: 'get',
    params
  })
}

// @Tags ItemData
// @Summary 分页获取项目数据列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取项目数据列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /itemDta/getItemDataList [get]
export const getItemDataList = (params) => {
  return service({
    url: '/itemDta/getItemDataList',
    method: 'get',
    params
  })
}

export const getItemDataListByCompany = (params) => {
  return service({
    url: '/itemDta/getItemDataListC',
    method: 'get',
    params
  })
}
