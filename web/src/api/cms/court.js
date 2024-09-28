import service from '@/utils/request'

// @Tags Court
// @Summary 创建法院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Court true "创建法院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /court/createCourt [post]
export const createCourt = (data) => {
  return service({
    url: '/court/createCourt',
    method: 'post',
    data
  })
}

// @Tags Court
// @Summary 删除法院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Court true "删除法院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /court/deleteCourt [delete]
export const deleteCourt = (params) => {
  return service({
    url: '/court/deleteCourt',
    method: 'delete',
    params
  })
}

// @Tags Court
// @Summary 批量删除法院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除法院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /court/deleteCourt [delete]
export const deleteCourtByIds = (params) => {
  return service({
    url: '/court/deleteCourtByIds',
    method: 'delete',
    params
  })
}

// @Tags Court
// @Summary 更新法院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Court true "更新法院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /court/updateCourt [put]
export const updateCourt = (data) => {
  return service({
    url: '/court/updateCourt',
    method: 'put',
    data
  })
}

// @Tags Court
// @Summary 用id查询法院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Court true "用id查询法院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /court/findCourt [get]
export const findCourt = (params) => {
  return service({
    url: '/court/findCourt',
    method: 'get',
    params
  })
}

// @Tags Court
// @Summary 分页获取法院列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取法院列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /court/getCourtList [get]
export const getCourtList = (params) => {
  return service({
    url: '/court/getCourtList',
    method: 'get',
    params
  })
}
