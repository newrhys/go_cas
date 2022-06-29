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

type CategoryApi struct {}

// @Tags Category
// @Summary 添加文章分类
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.ReqCategory true "分类名称,状态,排序"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"添加文章分类成功！"}"
// @Router /api/v1/category/category [post]
func (a *CategoryApi)AddCategory(ctx *gin.Context) {
	category := &request.ReqCategory{}
	if err := ctx.ShouldBindJSON(category);err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err := categoryService.AddCategory(*category); err != nil {
		global.GnLog.Error("添加文章分类失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.SuccessWithMessage(ctx, "添加文章分类成功！")
	}
}

// @Tags Category
// @Summary 编辑文章分类
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.ReqCategory true "分类名称,状态,排序"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"编辑文章分类成功！"}"
// @Router /api/v1/category/category/:id [put]
func (a *CategoryApi) UpdateCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	category := &request.ReqCategory{}
	if err := ctx.ShouldBindJSON(category);err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err := categoryService.UpdateCategory(*category, id); err != nil {
		global.GnLog.Error("编辑文章分类失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.SuccessWithMessage(ctx, "编辑文章分类成功！")
	}
}

// @Tags Category
// @Summary 删除文章分类
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"删除文章分类成功！"}"
// @Router /api/v1/category/category/:id [delete]
func (a *CategoryApi) DeleteCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := categoryService.DeleteCategory(id); err != nil {
		response.FailWithMessage(ctx, err.Error())
	} else {
		global.GnLog.Error("删除文章分类失败!", zap.Error(err))
		response.SuccessWithMessage(ctx, "删除文章分类成功！")
	}
}

// @Tags Category
// @Summary 文章分类列表
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.CategoryList true "分类名称"
// @Success 200 {string} string "{"code":200,"data":[],"msg":"获取文章分类列表成功！"}"
// @Router /api/v1/category/getCategoryList [get]
func (a *CategoryApi)CategoryList(ctx *gin.Context) {
	pageInfo := &request.CategoryList{}
	if err := ctx.ShouldBindQuery(pageInfo);err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err, list := categoryService.GetCategoryInfoList(*pageInfo); err != nil {
		global.GnLog.Error("获取文章分类列表失败!", zap.Error(err))
		response.FailWithMessage(ctx, "获取文章分类列表失败" + err.Error())
	} else {
		response.Success(ctx, list, "获取文章分类列表成功！")
	}
}

// @Tags Category
// @Summary 父级文章分类列表
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Success 200 {string} string "{"code":200,"data":[],"msg":"获取父级文章分类列表成功！"}"
// @Router /api/v1/category/parent [get]
func (a *CategoryApi)ParentList(ctx *gin.Context) {
	if list, err := categoryService.GetParentList(); err != nil {
		global.GnLog.Error("获取父级文章分类列表失败!", zap.Error(err))
		response.FailWithMessage(ctx, "获取父级文章分类列表失败" + err.Error())
	} else {
		response.Success(ctx, list, "获取父级文章分类列表成功！")
	}
}

// @Tags Category
// @Summary 子分类sort最大值
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Success 200 {string} string "{"code":200,"data":[],"msg":"获得子分类sort最大值成功！"}"
// @Router /api/v1/category/getSort/:id [get]
func (a *CategoryApi) GetSort(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if category,err := categoryService.GetSort(id); err != nil {
		global.GnLog.Error("获取子分类sort最大值失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.Success(ctx, category, "获得子分类sort最大值成功！")
	}
}
