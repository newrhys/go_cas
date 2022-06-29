package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
	"wave-admin/global"
	response2 "wave-admin/model/common/response"
	system2 "wave-admin/model/system"
	"wave-admin/model/system/request"
	"wave-admin/service"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

type JWT struct {
	SecretKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GnConfig.JWT.SecretKey),
	}
}

// jwt auth 认证
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtService := service.ServiceGroupApp.SystemServiceGroup.JwtService
		// 前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response2.Fail(ctx, 401, "权限不足！")
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		if jwtService.IsBlacklist(tokenString) {
			response2.Fail(ctx, 401, "您的帐户异地登陆或令牌失效")
			ctx.Abort()
			return
		}
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(tokenString)
		if err != nil {
			if err == TokenExpired {
				response2.Fail(ctx, 409, "授权已过期")
				ctx.Abort()
				return
			}
			response2.Fail(ctx, 401, "权限不足！")
			ctx.Abort()
			return
		}
		//log.Println("claims.ExpiresAt=", claims.ExpiresAt)
		//log.Println("time.Now=", time.Now().Unix())
		//log.Println("claims.BufferTime=", claims.BufferTime)
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.GnConfig.JWT.ExpiresTime
			newToken, _ := j.CreateToken(*claims)
			newClaims, _ := j.ParseToken(newToken)
			ctx.Header("New-Token", newToken)
			ctx.Header("New-Expires-At", strconv.FormatInt(newClaims.ExpiresAt, 10))
			if global.GnConfig.System.UseMultipoint {
				err, RedisJwtToken := jwtService.GetRedisJWT(newClaims.ID)
				if err != nil {
					global.GnLog.Error("get redis jwt failed", zap.Any("err", err))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = jwtService.JsonInBlacklist(system2.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.ID)
			}
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}

// 创建一个token
func (j *JWT) CreateToken(claims request.LoginClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SecretKey)
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*request.LoginClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.LoginClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SecretKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.LoginClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}