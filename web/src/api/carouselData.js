import service from '@/utils/request'

// @Tags CarouselData
// @Summary 创建轮播图数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CarouselData true "创建轮播图数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /caroDta/createCarouselData [post]
export const createCarouselData = (data) => {
  return service({
    url: '/caroDta/createCarouselData',
    method: 'post',
    data
  })
}

// @Tags CarouselData
// @Summary 删除轮播图数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CarouselData true "删除轮播图数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /caroDta/deleteCarouselData [delete]
export const deleteCarouselData = (data) => {
  return service({
    url: '/caroDta/deleteCarouselData',
    method: 'delete',
    data
  })
}

// @Tags CarouselData
// @Summary 批量删除轮播图数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除轮播图数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /caroDta/deleteCarouselData [delete]
export const deleteCarouselDataByIds = (data) => {
  return service({
    url: '/caroDta/deleteCarouselDataByIds',
    method: 'delete',
    data
  })
}

// @Tags CarouselData
// @Summary 更新轮播图数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CarouselData true "更新轮播图数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /caroDta/updateCarouselData [put]
export const updateCarouselData = (data) => {
  return service({
    url: '/caroDta/updateCarouselData',
    method: 'put',
    data
  })
}

// @Tags CarouselData
// @Summary 用id查询轮播图数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CarouselData true "用id查询轮播图数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /caroDta/findCarouselData [get]
export const findCarouselData = (params) => {
  return service({
    url: '/caroDta/findCarouselData',
    method: 'get',
    params
  })
}

// @Tags CarouselData
// @Summary 分页获取轮播图数据列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取轮播图数据列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /caroDta/getCarouselDataList [get]
export const getCarouselDataList = (params) => {
  return service({
    url: '/caroDta/getCarouselDataList',
    method: 'get',
    params
  })
}

export const getCarouselDataListByCompany = (params) => {
  return service({
    url: '/caroDta/getCarouselDataListC',
    method: 'get',
    params
  })
}
