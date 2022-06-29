package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"wave-admin/global"
	"wave-admin/model/common/response"
	"wave-admin/model/system/request"
	"wave-admin/utils"
)

type RecordApi struct{}

// @Tags Record
// @Summary 删除操作记录
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"删除操作记录成功！"}"
// @Router /api/v1/record/record/:id [delete]
func (a *RecordApi) DeleteRecord(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := recordService.DeleteRecord(id); err != nil {
		global.GnLog.Error("删除操作记录失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.SuccessWithMessage(ctx, "删除操作记录成功！")
	}
}

// @Tags Record
// @Summary 获取操作记录列表
// @Produce application/json
// @Success 200 {string} string "{"code":200,"data":{"list":[],"total":28,"current_page":1,"page_size":20},"msg":"获取操作记录成功！"}"
// @Router /api/v1/record/getRecordList [get]
func (a *RecordApi) RecordList(ctx *gin.Context)  {
	pageInfo := &request.RecordList{}
	if err := ctx.ShouldBindQuery(pageInfo);err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err, list, total := recordService.GetRecordInfoList(*pageInfo); err != nil {
		global.GnLog.Error("获取操作记录失败!", zap.Error(err))
		response.FailWithMessage(ctx, "获取操作记录失败" + err.Error())
	} else {
		response.Success(ctx, response.PageResult{
			List:        list,
			Total:       total,
			CurrentPage: pageInfo.CurrentPage,
			PageSize:    pageInfo.PageSize,
		}, "获取操作记录成功！")
	}
}
