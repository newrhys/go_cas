package global

import (
	"github.com/casbin/casbin/v2"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"wave-admin/cache/redis_key"
	"wave-admin/config"
)

var (
	GnVp      	*viper.Viper
	GnConfig   	config.Server
	GnLog      	*zap.Logger
	GnDb       	*gorm.DB
	GnRedis    	*redis.Client
	GnValidate 	*validator.Validate
	GnTrans    	ut.Translator
	GnRedisKey 	redis_key.RedisKey
	GnCasbin 	*casbin.SyncedEnforcer
)