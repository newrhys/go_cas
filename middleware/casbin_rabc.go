package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wave-admin/global"
	"wave-admin/model/common/response"
	"wave-admin/utils"
)

// 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleId := utils.GetRoleId(c)
		// 获取请求的PATH
		obj,_ := utils.GetRequestPath(c.Request.URL.Path)

		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := strconv.Itoa(int(roleId))
		e := global.GnCasbin
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if success {
			c.Next()
		} else {
			response.FailWithMessage(c, "权限不足！")
			c.Abort()
			return
		}
	}
}
