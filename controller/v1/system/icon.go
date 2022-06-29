package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"wave-admin/global"
	"wave-admin/model/common/response"
	"wave-admin/model/system/request"
	"wave-admin/utils"
)

type IconApi struct{}

// @Tags Icon
// @Summary 图标列表
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.IconList true "名称"
// @Success 200 {string} string "{"code":200,"data":[],"msg":"获取图标列表成功！"}"
// @Router /api/v1/icon/getIconList [get]
func (a *IconApi) IconList(ctx *gin.Context) {
	pageInfo := &request.IconList{}
	if err := ctx.ShouldBindQuery(pageInfo);err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err, list := iconService.GetIconInfoList(*pageInfo); err != nil {
		global.GnLog.Error("获取图标列表失败!", zap.Error(err))
		response.FailWithMessage(ctx, "获取图标列表失败" + err.Error())
	} else {
		response.Success(ctx, list, "获取图标列表成功！")
	}
}
