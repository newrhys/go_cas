package cms

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"wave-admin/global"
	"wave-admin/model/cms/request"
	"wave-admin/model/common/response"
	"wave-admin/utils"
)

type TagApi struct {}

// @Tags Tag
// @Summary 添加标签
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.ReqTag true "标签名称,状态,排序"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"添加标签成功！"}"
// @Router /api/v1/tag/tag [post]
func (a *TagApi)AddTag(ctx *gin.Context) {
	tag := &request.ReqTag{}
	if err := ctx.ShouldBindJSON(tag);err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err := tagService.AddTag(*tag); err != nil {
		global.GnLog.Error("添加标签失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.SuccessWithMessage(ctx, "添加标签成功！")
	}
}

// @Tags Tag
// @Summary 编辑标签
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.ReqTag true "标签名称,状态,排序"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"编辑标签成功！"}"
// @Router /api/v1/tag/tag/:id [put]
func (a *TagApi) UpdateTag(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	tag := &request.ReqTag{}
	if err := ctx.ShouldBindJSON(tag);err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err := tagService.UpdateTag(*tag, id); err != nil {
		global.GnLog.Error("编辑标签失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.SuccessWithMessage(ctx, "编辑标签成功！")
	}
}

// @Tags Tag
// @Summary 删除标签
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"删除标签成功！"}"
// @Router /api/v1/tag/tag/:id [delete]
func (a *TagApi) DeleteTag(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := tagService.DeleteTag(id); err != nil {
		response.FailWithMessage(ctx, err.Error())
	} else {
		global.GnLog.Error("删除标签失败!", zap.Error(err))
		response.SuccessWithMessage(ctx, "删除标签成功！")
	}
}

// @Tags Tag
// @Summary 标签列表
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.TagList true "标签名称,状态"
// @Success 200 {string} string "{"code":200,"data":{"list":[],"total":7,"current_page":1,"page_size":20},"msg":"获取标签列表成功！"}"
// @Router /api/v1/tag/getTagList [get]
func (a *TagApi)TagList(ctx *gin.Context) {
	pageInfo := &request.TagList{}
	if err := ctx.ShouldBindQuery(pageInfo);err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err, list, total := tagService.GetTagList(*pageInfo); err != nil {
		global.GnLog.Error("获取标签列表失败!", zap.Error(err))
		response.FailWithMessage(ctx, "获取标签列表失败" + err.Error())
	} else {
		response.Success(ctx, response.PageResult{
			List:        list,
			Total:       total,
			CurrentPage: pageInfo.CurrentPage,
			PageSize:    pageInfo.PageSize,
		}, "获取标签列表成功！")
	}
}

// @Tags Tag
// @Summary 标签sort最大值
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"获得标签sort最大值成功！"}"
// @Router /api/v1/tag/getSort [get]
func (a *TagApi) GetSort(ctx *gin.Context) {
	if tag,err := tagService.GetSort(); err != nil {
		global.GnLog.Error("获取标签sort最大值失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.Success(ctx, tag, "获得标签sort最大值成功！")
	}
}

// @Tags Tag
// @Summary 选择标签列表
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"获得选择标签列表成功！"}"
// @Router /api/v1/tag/getSelectTagList [get]
func (a *TagApi) SelectTagList(ctx *gin.Context) {
	if err, list := tagService.SelectTagList(); err != nil {
		global.GnLog.Error("获取选择标签列表失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.Success(ctx, list, "获得选择标签列表成功！")
	}
}