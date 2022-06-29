package initialize

import (
	"fmt"
	"os"
	"wave-admin/global"

	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.GnConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.GnLog.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.GnLog.Info("redis connect ping response:", zap.String("pong", pong))
		global.GnRedis = client
	}

	// 将redis key 放到全局变量
	workDir, _ := os.Getwd()
	v := viper.New()
	v.SetConfigName("redis_key")
	v.SetConfigType("yml")
	// v.AddConfigPath(workDir + "/cache/redis_key")
	v.AddConfigPath(workDir)

	err2 := v.ReadInConfig()
	if err2 != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err2))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GnRedisKey); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GnRedisKey); err != nil {
		fmt.Println(err)
	}
}
