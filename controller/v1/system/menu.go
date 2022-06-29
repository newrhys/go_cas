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

type MenuApi struct{}

func (a *MenuApi) GetMenu(ctx *gin.Context)  {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if authority,err := menuService.GetMenu(id); err != nil {
		global.GnLog.Error("获取菜单失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.Success(ctx, authority, "获得菜单成功！")
	}
}

// @Tags Menu
// @Summary 添加菜单
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.ReqMenu true "菜单名称,权限code,排序"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"添加菜单成功！"}"
// @Router /api/v1/menu/menu [post]
func (a *MenuApi) AddMenu(ctx *gin.Context) {
	menu := &request.ReqMenu{}
	if err := ctx.ShouldBindJSON(menu);err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err := menuService.AddMenu(*menu); err != nil {
		global.GnLog.Error("添加菜单失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.SuccessWithMessage(ctx, "添加菜单成功！")
	}
}

// @Tags Menu
// @Summary 编辑菜单
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.ReqMenu true "菜单名称,权限code,排序"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"编辑菜单成功！"}"
// @Router /api/v1/menu/menu/:id [put]
func (a *MenuApi) UpdateMenu(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	menu := &request.ReqMenu{}
	if err := ctx.ShouldBindJSON(menu);err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err := menuService.UpdateMenu(*menu, id); err != nil {
		global.GnLog.Error("编辑菜单失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.SuccessWithMessage(ctx, "编辑菜单成功！")
	}
}

// @Tags Menu
// @Summary 删除菜单
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"删除菜单成功！"}"
// @Router /api/v1/menu/menu/:id [delete]
func (a *MenuApi) DeleteMenu(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := menuService.DeleteMenu(id); err != nil {
		response.FailWithMessage(ctx, err.Error())
	} else {
		global.GnLog.Error("删除菜单失败!", zap.Error(err))
		response.SuccessWithMessage(ctx, "删除菜单成功！")
	}
}

// @Tags Menu
// @Summary 菜单列表
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.MenuList true "描述"
// @Success 200 {string} string "{"code":200,"data":[],"msg":"获取菜单列表成功！"}"
// @Router /api/v1/menu/getMenuList [get]
func (a *MenuApi) MenuList(ctx *gin.Context) {
	pageInfo := &request.MenuList{}
	if err := ctx.ShouldBindQuery(pageInfo);err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err, list := menuService.GetMenuInfoList(*pageInfo); err != nil {
		global.GnLog.Error("获取菜单列表失败!", zap.Error(err))
		response.FailWithMessage(ctx, "获取菜单列表失败" + err.Error())
	} else {
		response.Success(ctx, list, "获取菜单列表成功！")
	}
}

// @Tags Menu
// @Summary 父级菜单列表
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Success 200 {string} string "{"code":200,"data":[],"msg":"获取父级菜单列表成功！"}"
// @Router /api/v1/menu/parent [get]
func (a *MenuApi) ParentList(ctx *gin.Context)  {
	if list, err := menuService.GetParentList(); err != nil {
		global.GnLog.Error("获取父级菜单列表失败!", zap.Error(err))
		response.FailWithMessage(ctx, "获取父级菜单列表失败" + err.Error())
	} else {
		response.Success(ctx, list, "获取父级菜单列表成功！")
	}
}

// @Tags Menu
// @Summary 子菜单sort最大值
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Success 200 {string} string "{"code":200,"data":[],"msg":"获得子菜单sort最大值成功！"}"
// @Router /api/v1/menu/getSort/:id [get]
func (a *MenuApi) GetSort(ctx *gin.Context)  {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if menu,err := menuService.GetSort(id); err != nil {
		global.GnLog.Error("获取子菜单sort最大值失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.Success(ctx, menu, "获得子菜单sort最大值成功！")
	}
}