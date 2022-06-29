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

type SysApiApi struct{}

func (a *SysApiApi) GetApi(ctx *gin.Context)  {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if role,err := apiService.GetApi(id); err != nil {
		global.GnLog.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.Success(ctx, role, "获得API成功！")
	}
}

// @Tags Api
// @Summary 添加API
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.ReqApi true "描述"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"添加API成功！"}"
// @Router /api/v1/api/api [post]
func (a *SysApiApi) AddApi(ctx *gin.Context) {
	api := &request.ReqApi{}
	if err := ctx.ShouldBindJSON(api);err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err := apiService.AddApi(*api); err != nil {
		global.GnLog.Error("添加API失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.SuccessWithMessage(ctx, "添加API成功！")
	}
}

// @Tags Api
// @Summary 编辑API
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.ReqApi true "描述"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"编辑API成功！"}"
// @Router /api/v1/api/api/:id [put]
func (a *SysApiApi) UpdateApi(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	api := &request.ReqApi{}
	if err := ctx.ShouldBindJSON(api);err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err := apiService.UpdateApi(*api, id); err != nil {
		global.GnLog.Error("编辑API失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.SuccessWithMessage(ctx, "编辑API成功！")
	}
}

// @Tags Api
// @Summary 删除API
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"删除API成功！"}"
// @Router /api/v1/api/api/:id [delete]
func (a *SysApiApi) DeleteApi(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := apiService.DeleteApi(id); err != nil {
		global.GnLog.Error("删除API失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.SuccessWithMessage(ctx, "删除API成功！")
	}
}

// @Tags Api
// @Summary API列表
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.ApiList true "菜单名称,状态"
// @Success 200 {string} string "{"code":200,"data":{"list":[],"total":0,"current_page":1,"page_size":20},"msg":"获取API列表成功！"}"
// @Router /api/v1/api/getApiList [get]
func (a *SysApiApi) ApiList(ctx *gin.Context) {
	pageInfo := &request.ApiList{}
	if err := ctx.ShouldBindQuery(pageInfo);err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err, list, total := apiService.GetApiInfoList(*pageInfo); err != nil {
		global.GnLog.Error("获取API列表失败!", zap.Error(err))
		response.FailWithMessage(ctx, "获取API列表失败" + err.Error())
	} else {
		response.Success(ctx, response.PageResult{
			List:        list,
			Total:       total,
			CurrentPage: pageInfo.CurrentPage,
			PageSize:    pageInfo.PageSize,
		}, "获取API列表成功！")
	}
}

// @Tags Api
// @Summary 父级API列表
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Success 200 {string} string "{"code":200,"data":[],"msg":"获取父级API列表成功！"}"
// @Router /api/v1/api/parent [get]
func (a *SysApiApi) ParentList(ctx *gin.Context)  {
	if list, err := apiService.GetTreeList(); err != nil {
		global.GnLog.Error("获取父级API列表失败!", zap.Error(err))
		response.FailWithMessage(ctx, "获取父级API列表失败" + err.Error())
	} else {
		response.Success(ctx, list, "获取父级API列表成功！")
	}
}
