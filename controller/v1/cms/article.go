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

type ArticleApi struct {}

// @Tags Article
// @Summary 文章内容
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"获取文章成功！"}"
// @Router /api/v1/article/article/:id [get]
func (a *ArticleApi)GetArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err, article := articleService.GetArticle(id); err != nil {
		global.GnLog.Error("获取文章失败!", zap.Error(err))
		response.FailWithMessage(ctx, "获取文章失败" + err.Error())
	} else {
		response.Success(ctx, article, "获取文章成功！")
	}
}

// @Tags Article
// @Summary 添加文章
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.ReqArticle true "文章标题"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"添加文章成功！"}"
// @Router /api/v1/article/article [post]
func (a *ArticleApi)AddArticle(ctx *gin.Context) {
	article := &request.ReqArticle{}
	if err := ctx.ShouldBindJSON(article);err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	userId := utils.GetUserID(ctx)
	if err := articleService.AddArticle(userId, *article); err != nil {
		global.GnLog.Error("添加文章失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.SuccessWithMessage(ctx, "添加文章成功！")
	}
}

// @Tags Article
// @Summary 编辑文章
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.ReqArticle true "文章标题"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"编辑文章成功！"}"
// @Router /api/v1/article/article/:id [put]
func (a *ArticleApi) UpdateArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	article := &request.ReqArticle{}
	if err := ctx.ShouldBindJSON(article);err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	userId := utils.GetUserID(ctx)
	if err := articleService.UpdateArticle(userId, *article, id); err != nil {
		global.GnLog.Error("编辑文章失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.SuccessWithMessage(ctx, "编辑文章成功！")
	}
}

// @Tags Article
// @Summary 删除文章
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"删除文章成功！"}"
// @Router /api/v1/article/article/:id [delete]
func (a *ArticleApi) DeleteArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := articleService.DeleteArticle(id); err != nil {
		response.FailWithMessage(ctx, err.Error())
	} else {
		global.GnLog.Error("删除文章失败!", zap.Error(err))
		response.SuccessWithMessage(ctx, "删除文章成功！")
	}
}

// @Tags Article
// @Summary 文章列表
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.ArticleList true "文章标题,状态"
// @Success 200 {string} string "{"code":200,"data":{"list":[],"total":7,"current_page":1,"page_size":20},"msg":"获取文章列表成功！"}"
// @Router /api/v1/article/getArticleList [get]
func (a *ArticleApi)ArticleList(ctx *gin.Context) {
	pageInfo := &request.ArticleList{}
	if err := ctx.ShouldBindQuery(pageInfo);err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err, list, total := articleService.GetArticleList(*pageInfo); err != nil {
		global.GnLog.Error("获取文章列表失败!", zap.Error(err))
		response.FailWithMessage(ctx, "获取文章列表失败" + err.Error())
	} else {
		response.Success(ctx, response.PageResult{
			List:        list,
			Total:       total,
			CurrentPage: pageInfo.CurrentPage,
			PageSize:    pageInfo.PageSize,
		}, "获取文章列表成功！")
	}
}
