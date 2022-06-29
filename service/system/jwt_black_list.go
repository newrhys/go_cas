package system

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
	"time"
	"wave-admin/global"
	"wave-admin/model/system"
)

type JwtService struct{}

// 拉黑jwt
func (jwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.GnDb.Create(&jwtList).Error
	return
}

// 判断JWT是否在黑名单内部
func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	isNotFound := errors.Is(global.GnDb.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error, gorm.ErrRecordNotFound)
	return !isNotFound
}

// 从redis取jwt
func (jwtService *JwtService) GetRedisJWT(userId uint64) (err error, redisJWT string) {
	key := global.GnRedisKey.Jwt.JwtKey + strconv.FormatUint(userId,20)
	redisJWT, err = global.GnRedis.Get(key).Result()
	return err, redisJWT
}

// jwt存入redis并设置过期时间
func (jwtService *JwtService) SetRedisJWT(jwt string, userId uint64) (err error) {
	key := global.GnRedisKey.Jwt.JwtKey + strconv.FormatUint(userId,20)
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.GnConfig.JWT.ExpiresTime) * time.Second
	err = global.GnRedis.Set(key, jwt, timer).Err()
	return err
}

// 删除redis存的jwt
func (jwtService *JwtService) DelRedisJWT(userId uint64)  (err error) {
	key := global.GnRedisKey.Jwt.JwtKey + strconv.FormatUint(userId,20)
	err = global.GnRedis.Del(key).Err()
	return err
}