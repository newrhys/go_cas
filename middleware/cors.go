package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		//origin := ctx.Request.Header.Get("Origin")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Language, Content-Type, Expires, Last-Modified, New-Token, New-Expires-At")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		ctx.Next()
	}
}