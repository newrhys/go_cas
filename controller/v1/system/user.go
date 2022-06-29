package system

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"wave-admin/global"
	"wave-admin/middleware"
	"wave-admin/model/common/response"
	"wave-admin/model/system"
	"wave-admin/model/system/request"
	SystemRep "wave-admin/model/system/response"
	"wave-admin/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type UserApi struct{}

var store = base64Captcha.DefaultMemStore

// @Tags Base
// @Summary 验证码
// @Produce application/json
// @Success 200 {string} string "{"code":200,"data":{"captcha_id":"captcha_id","captcha_length":5},"msg":"获得验证码成功！"}"
// @Router /api/v1/auth/captcha [get]
func (a *UserApi) Captcha(ctx *gin.Context) {
	// 字符，公式验证码配置
	driver := base64Captcha.NewDriverDigit(global.GnConfig.Captcha.ImgHeight, global.GnConfig.Captcha.ImgWidth, global.GnConfig.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.GnLog.Error("验证码获取失败！", zap.Error(err))
		response.FailWithMessage(ctx, "验证码获取失败！")
	} else {
		response.Success(ctx, SystemRep.CaptchaResponse{
			CaptchaId:     id,
			PicPath:       b64s,
			CaptchaLength: global.GnConfig.Captcha.KeyLong,
		}, "获得验证码成功！")
	}
}

// @Tags Base
// @Summary 登录
// @Produce application/json
// @Param data body request.Login true "用户名,密码"
// @Success 200 {string} string "{"code":200,"data":{"token":"token","expiresAt":1656042399},"msg":"登录成功！"}"
// @Router /api/v1/auth/login [post]
func (a *UserApi) Login(ctx *gin.Context) {
	login := &request.Login{}
	if err := ctx.ShouldBindJSON(login); err != nil {
		global.GnLog.Error("login error: ", zap.Any("err", err))
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	//flag := userService.Verify(login.CaptchaId, login.VerifyCode)
	//if flag {
	user, err := userService.Login(*login)
	if err != nil {
		response.Fail(ctx, 422, err.Error())
		return
	}

	// 发放token
	issueToken(ctx, user)
	//} else {
	//	response.FailWithMessage(ctx, "验证码错误！")
	//}
}

// 发放token
func issueToken(ctx *gin.Context, user system.SysUser) {
	log.Println(user)
	var userRole system.SysUserRole
	err := global.GnDb.Where("user_id = ?", user.ID).Find(&userRole).Error
	log.Println(userRole)
	if err != nil {
		response.Fail(ctx, 500, "系统异常！")
		return
	}
	if userRole.RoleId == 0 {
		response.FailWithMessage(ctx, "该用户未分角色组！")
		return
	}
	j := &middleware.JWT{SecretKey: []byte(global.GnConfig.JWT.SecretKey)} // 唯一签名
	claims := request.LoginClaims{
		ID:         user.ID,
		RoleId:     userRole.RoleId,
		Username:   user.Username,
		BufferTime: global.GnConfig.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                            // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.GnConfig.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "gin demo",
			Subject:   "user token",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		response.Fail(ctx, 500, "系统异常！")
		global.GnLog.Error("token generate error: ", zap.Any("err", err))
		return
	}
	repUser := SystemRep.LoginResponse{
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt,
	}
	if !global.GnConfig.System.UseMultipoint {
		response.Success(ctx, repUser, "成功！")
		return
	}
	if err, jwtStr := jwtService.GetRedisJWT(user.ID); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.ID); err != nil {
			global.GnLog.Error("设置登录状态失败", zap.Any("err", err))
			response.Fail(ctx, 422, "设置登录状态失败1")
			return
		}
		response.Success(ctx, repUser, "成功！")
	} else if err != nil {
		global.GnLog.Error("设置登录状态失败", zap.Any("err", err))
		response.Fail(ctx, 422, "设置登录状态失败2")
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.Fail(ctx, 422, "jwt作废失败")
			return
		}
		if err := jwtService.SetRedisJWT(token, user.ID); err != nil {
			response.Fail(ctx, 422, "设置登录状态失败3")
			return
		}
		response.Success(ctx, repUser, "登录成功！")
	}
}

// @Tags User
// @Summary 退出
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"退出成功！"}"
// @Router /api/v1/auth/loginOut [post]
func (a *UserApi) LoginOut(ctx *gin.Context) {
	userId := utils.GetUserID(ctx)
	if err := jwtService.DelRedisJWT(userId); err != nil {
		response.FailWithMessage(ctx, "退出登录失败！")
		return
	}
	response.SuccessWithMessage(ctx, "退出成功！")
}

// @Tags User
// @Summary 添加系统用户
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.AddUser true "用户名,密码,昵称,状态"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"添加用户成功！"}"
// @Router /api/v1/user/user [post]
func (a *UserApi) AddUser(ctx *gin.Context) {
	user := &request.AddUser{}
	if err := ctx.ShouldBindJSON(user); err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err := userService.AddUser(*user); err != nil {
		global.GnLog.Error("添加用户失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.SuccessWithMessage(ctx, "添加用户成功！")
	}
}

// @Tags User
// @Summary 登录用户基本信息
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Success 200 {string} string "{"code":200,"data":{"id":1,"uuid":"uuid","username":"username","nickname":"nickname","avatar":"avatar","status":1,"mobile":"","email":"","roles":[]},"msg":"获得登录用户信息成功！"}"
// @Router /api/v1/user/info [get]
func (a *UserApi) GetUserInfo(ctx *gin.Context) {
	id := utils.GetUserID(ctx)
	fmt.Printf("id: %v\n", id)
	if user, err := userService.GetUser(id); err != nil {
		global.GnLog.Error("获取登录用户信息失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.Success(ctx, user, "获得登录用户信息成功！")
	}
}

// @Tags User
// @Summary 用户Route菜单
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Success 200 {string} string "{"code":200,"data":{"route_tree":[]}]}]},"msg":"获得Route菜单成功！"}"
// @Router /api/v1/user/getRouteMenuList [get]
func (a *UserApi) GetRouteMenuList(ctx *gin.Context) {
	id := utils.GetUserID(ctx)
	if list, err := userService.GetRouteMenuList(id); err != nil {
		global.GnLog.Error("获取Route菜单失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.Success(ctx, list, "获得Route菜单成功！")
	}
}

func (a *UserApi) GetUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if user, err := userService.GetUser(uint64(id)); err != nil {
		global.GnLog.Error("获取用户信息失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.Success(ctx, user, "获得用户信息成功！")
	}
}

// @Tags User
// @Summary 编辑用户
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.UpdateUser true "用户名,昵称"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"编辑用户成功！"}"
// @Router /api/v1/user/user/:id [put]
func (a *UserApi) UpdateUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user := &request.UpdateUser{}
	if err := ctx.ShouldBindJSON(user); err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err := userService.UpdateUser(*user, id); err != nil {
		global.GnLog.Error("编辑用户失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.SuccessWithMessage(ctx, "编辑用户成功！")
	}
}

// @Tags User
// @Summary 删除用户
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"删除用户成功！"}"
// @Router /api/v1/user/user/:id [delete]
func (a *UserApi) DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := userService.DeleteUser(id); err != nil {
		global.GnLog.Error("删除用户失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.SuccessWithMessage(ctx, "删除用户成功！")
	}
}

// @Tags User
// @Summary 系统用户列表
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.UserList true "用户名"
// @Success 200 {string} string "{"code":200,"data":{"list":[],"total":4,"current_page":1,"page_size":20},"msg":"获取用户列表成功！"}"
// @Router /api/v1/user/getUserList [get]
func (a *UserApi) GetUserList(ctx *gin.Context) {
	pageInfo := &request.UserList{}
	if err := ctx.ShouldBindQuery(pageInfo); err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err, list, total := userService.GetUserList(*pageInfo); err != nil {
		global.GnLog.Error("获取用户列表失败!", zap.Error(err))
		response.FailWithMessage(ctx, "获取用户列表失败！")
	} else {
		response.Success(ctx, response.PageResult{
			List:        list,
			Total:       total,
			CurrentPage: pageInfo.CurrentPage,
			PageSize:    pageInfo.PageSize,
		}, "获取用户列表成功！")
	}
}

// @Tags User
// @Summary 获得用户角色id
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Success 200 {string} string "{"code":200,"data":{"id":1,"created_at":"created_at","user_id":1,"role_id":1},"msg":"获得用户角色id成功！"}"
// @Router /api/v1/user/getRoleIdByUserId/:id [get]
func (a *UserApi) GetRoleIdByUserId(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if role, err := userService.GetRoleIdByUserId(id); err != nil {
		global.GnLog.Error("获得用户角色id失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.Success(ctx, role, "获得用户角色id成功！")
	}
}

// @Tags User
// @Summary 分配角色
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.AssignRole true "用户id, 角色id"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"分配角色成功！"}"
// @Router /api/v1/user/getRoleIdByUserId/:id [get]
func (a *UserApi) AssignRole(ctx *gin.Context) {
	assignRole := &request.AssignRole{}
	if err := ctx.ShouldBindJSON(assignRole); err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err := userService.AssignRole(*assignRole); err != nil {
		global.GnLog.Error("分配角色失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.SuccessWithMessage(ctx, "分配角色成功！")
	}
}

// @Tags User
// @Summary 修改密码
// @Produce application/json
// @Param Authorization header string true "验证参数Bearer和token空格拼接"
// @Param data body request.ChangePassword true "旧密码,新密码,确认密码"
// @Success 200 {string} string "{"code":200,"data":null,"msg":"修改密码成功！"}"
// @Router /api/v1/user/changepassword [put]
func (a *UserApi) ChangePassword(ctx *gin.Context) {
	id := utils.GetUserID(ctx)
	pwd := &request.ChangePassword{}
	if err := ctx.ShouldBindJSON(pwd); err != nil {
		response.FailWithMessage(ctx, utils.Error(err))
		return
	}
	if err := userService.ChangePassword(*pwd, id); err != nil {
		global.GnLog.Error("修改密码失败!", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.SuccessWithMessage(ctx, "修改密码成功！")
	}
}
