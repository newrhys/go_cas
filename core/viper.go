package core

import (
	"fmt"
	"os"
	"wave-admin/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitViper() *viper.Viper {
	workDir, _ := os.Getwd()
	v := viper.New()
	v.SetConfigName("application")
	v.SetConfigType("yml")
	// v.AddConfigPath(workDir + "/config")
	v.AddConfigPath(workDir)

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GnConfig); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GnConfig); err != nil {
		fmt.Println(err)
	}
	return v
}
