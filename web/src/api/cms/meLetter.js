import service from '@/utils/request'

// @Tags MeLetter
// @Summary 创建我的保函
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MeLetter true "创建我的保函"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /meLetter/createMeLetter [post]
export const createMeLetter = (data) => {
  return service({
    url: '/meLetter/createMeLetter',
    method: 'post',
    data
  })
}

// @Tags MeLetter
// @Summary 删除我的保函
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MeLetter true "删除我的保函"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /meLetter/deleteMeLetter [delete]
export const deleteMeLetter = (params) => {
  return service({
    url: '/meLetter/deleteMeLetter',
    method: 'delete',
    params
  })
}

// @Tags MeLetter
// @Summary 批量删除我的保函
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除我的保函"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /meLetter/deleteMeLetter [delete]
export const deleteMeLetterByIds = (params) => {
  return service({
    url: '/meLetter/deleteMeLetterByIds',
    method: 'delete',
    params
  })
}

// @Tags MeLetter
// @Summary 更新我的保函
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MeLetter true "更新我的保函"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /meLetter/updateMeLetter [put]
export const updateMeLetter = (data) => {
  return service({
    url: '/meLetter/updateMeLetter',
    method: 'put',
    data
  })
}

// @Tags MeLetter
// @Summary 用id查询我的保函
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MeLetter true "用id查询我的保函"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /meLetter/findMeLetter [get]
export const findMeLetter = (params) => {
  return service({
    url: '/meLetter/findMeLetter',
    method: 'get',
    params
  })
}

// @Tags MeLetter
// @Summary 分页获取我的保函列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取我的保函列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /meLetter/getMeLetterList [get]
export const getMeLetterList = (params) => {
  return service({
    url: '/meLetter/getMeLetterList',
    method: 'get',
    params
  })
}
