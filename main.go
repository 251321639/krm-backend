// 项目总入口
package main

import (
	"krm-backend/config"
	_ "krm-backend/controllers/initcontroller"
	"krm-backend/middlerwares"
	"krm-backend/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	//1.加载程序的配置
	//2.配置gin
	r := gin.Default()
	if config.RunAsMode == "debug" || config.RunAsMode == "dev" {
		r.Use(middlewares.Cors)
	}
	r.Use(middlewares.JWTAuth)
	//将路由传递给routers
	routers.RegisterRoutes(r)
	r.Run(config.Port)

}
