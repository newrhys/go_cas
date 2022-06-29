package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"wave-admin/global"
	"wave-admin/model/common/response"
	"wave-admin/model/system/request"
	SystemRep "wave-admin/model/system/response"
	"wave-admin/utils"
)

type RoleApi struct{}

func (a *RoleApi) GetRole(ctx *gin.Context)  {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if role,err := roleService.GetRole(id); err != nil {
		global.GnLog.Error("获取角色失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.Success(ctx, role, "获得角色成功！")
	}
}

// @Tags Role
// @Summary 获取角色菜单列表
// @Produce application/json
// @Success 200 {string} string "{"code":200,"data":{"check_list":[],"menu_list":[]},"msg":"获取角色菜单列表成功！"}"
// @Router /api/v1/role/getAssignPermissionTree/:id [get]
func (a *RoleApi) GetAssignPermissionTree(ctx *gin.Context)  {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if list,err := menuService.GetParentList(); err != nil {
		global.GnLog.Error("获取角色菜单列表失败!", zap.Error(err))
		response.FailWithMessage(ctx, "获取角色菜单列表失败" + err.Error())
	} else {
		if err, checks := roleService.GetAssignMenu(id); err != nil {
			global.GnLog.Error("获取角色菜单列表失败!", zap.Error(err))
			response.FailWithMessage(ctx, "获取父角色菜单列表失败" + err.Error())
		} else {
			response.Success(ctx, SystemRep.AssignRoleMenu{
				CheckList: checks,
				MenuList:  list,
			}, "获取角色菜单列表成功！")
		}
	}
}

// @Tags Role
// @Summary 获取角色API列表
// @Produce application/json
// @Success 200 {string} string "{"code":200,"data":{"check_list":[],"api_list":[]},"msg":"获取角色API列表成功！"}"
// @Router /api/v1/role/getAssignPermissionApi/:id [get]
func (a *RoleApi) GetAssignPermissionApi(ctx *gin.Context)  {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if list,err := apiService.GetTreeList(); err != nil {
		global.GnLog.Error("获取角色API列表失败!", zap.Error(err))
		response.FailWithMessage(ctx, "获取角色API列表失败" + err.Error())
	} else {
		if err, checks := roleService.GetAssignApi(id); err != nil {
			global.GnLog.Error("获取角色API列表失败!", zap.Error(err))
			response.FailWithMessage(ctx, "获取角色API列表失败" + err.Error())
		} else {
			response.Success(ctx, SystemRep.AssignRoleApi{
				CheckList: checks,
				ApiList:  list,
			}, "获取角色API列表成功！")
		}
	}
}

// @Tags Role
// @Summary 添加角色
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.ReqRole true "角色名称"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"添加角色成功！"}"
// @Router /api/v1/role/role [post]
func (a *RoleApi) AddRole(ctx *gin.Context) {
	role := &request.ReqRole{}
	if err := ctx.ShouldBindJSON(role);err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err,_ := roleService.AddRole(*role); err != nil {
		global.GnLog.Error("添加角色失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.SuccessWithMessage(ctx, "添加角色成功！")
	}
}

// @Tags Role
// @Summary 编辑角色
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.ReqRole true "角色名称"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"编辑角色成功！"}"
// @Router /api/v1/role/role/:id [put]
func (a *RoleApi) UpdateRole(ctx *gin.Context) {
	role := &request.ReqRole{}
	if err := ctx.ShouldBindJSON(role);err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err := roleService.UpdateRole(*role); err != nil {
		global.GnLog.Error("编辑角色失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.SuccessWithMessage(ctx, "编辑角色成功！")
	}
}

// @Tags Role
// @Summary 删除角色
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"删除角色成功！"}"
// @Router /api/v1/role/role/:id [delete]
func (a *RoleApi) DeleteRole(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := roleService.DeleteRole(id); err != nil {
		global.GnLog.Error("删除角色失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.SuccessWithMessage(ctx, "删除角色成功！")
	}
}

// @Tags Role
// @Summary 获取角色列表
// @Produce application/json
// @Success 200 {string} string "{"code":200,"data":{"check_list":[],"api_list":[]},"msg":"获取角色API列表成功！"}"
// @Router /api/v1/role/getAssignPermissionApi/:id [get]
func (a *RoleApi) RoleList(ctx *gin.Context)  {
	pageInfo := &request.RoleList{}
	if err := ctx.ShouldBindQuery(pageInfo);err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err, list, total := roleService.GetRoleInfoList(*pageInfo); err != nil {
		global.GnLog.Error("获取角色列表失败!", zap.Error(err))
		response.FailWithMessage(ctx, "获取角色列表失败" + err.Error())
	} else {
		response.Success(ctx, response.PageResult{
			List:        list,
			Total:       total,
			CurrentPage: pageInfo.CurrentPage,
			PageSize:    pageInfo.PageSize,
		}, "获取角色列表成功！")
	}
}

// @Tags Role
// @Summary 分配权限
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.AssignSave true "角色id"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"分配权限成功！"}"
// @Router /api/v1/role/roleAssignSave [post]
func (a *RoleApi) AssignSave(ctx *gin.Context)  {
	assignSave := &request.AssignSave{}
	if err := ctx.ShouldBindJSON(assignSave);err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err,casbinInfos := roleService.AssignSave(*assignSave); err != nil {
		global.GnLog.Error("分配权限失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		// 添加casbin
		_ = casbinService.UpdateCasbin(assignSave.ID, casbinInfos)
		response.SuccessWithMessage(ctx, "分配权限成功！")
	}
}