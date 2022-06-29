package core

import (
	"fmt"
	"wave-admin/global"
	"wave-admin/initialize"
)

func RunServer()  {
	if global.GnConfig.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}

	r := initialize.InitRouter()

	port := global.GnConfig.System.ServerPort
	fmt.Printf("默认自动化文档地址:http://127.0.0.1:%s/swagger/index.html\n", port)
	fmt.Printf("默认后端接口地址:http://127.0.0.1:%s\n", port)
	if port != "" {
		r.Run(":" + port)
	} else {
		r.Run()
	}
}