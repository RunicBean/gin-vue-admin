import service from '@/utils/request'

// @Tags Company
// @Summary 创建公司
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Company true "创建公司"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /company/createCompany [post]
export const createCompany = (data) => {
  return service({
    url: '/company/createCompany',
    method: 'post',
    data
  })
}

// @Tags Company
// @Summary 删除公司
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Company true "删除公司"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /company/deleteCompany [delete]
export const deleteCompany = (data) => {
  return service({
    url: '/company/deleteCompany',
    method: 'delete',
    data
  })
}

// @Tags Company
// @Summary 批量删除公司
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除公司"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /company/deleteCompany [delete]
export const deleteCompanyByIds = (data) => {
  return service({
    url: '/company/deleteCompanyByIds',
    method: 'delete',
    data
  })
}

// @Tags Company
// @Summary 更新公司
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Company true "更新公司"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /company/updateCompany [put]
export const updateCompany = (data) => {
  return service({
    url: '/company/updateCompany',
    method: 'put',
    data
  })
}

// @Tags Company
// @Summary 用id查询公司
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Company true "用id查询公司"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /company/findCompany [get]
export const findCompany = (params) => {
  return service({
    url: '/company/findCompany',
    method: 'get',
    params
  })
}

export const findCompanyByPath = (params) => {
  return service({
    url: '/company/findCName',
    method: 'get',
    params
  })
}

// @Tags Company
// @Summary 分页获取公司列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取公司列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /company/getCompanyList [get]
export const getCompanyList = (params) => {
  return service({
    url: '/company/getCompanyList',
    method: 'get',
    params
  })
}


