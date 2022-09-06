// @File    :   app.go
// @Time    :   2020/11/23 22:52:45
// @Author  :   Lison
// @Version :   1.0
// @Contact :   linliangsuan@qq.com
// @License :   (C)Copyright 2017-2018, Liugroup-NLPR-CASIA
// @Desc    :   初始化

package app

import (
	"gdopo.com/config"
	"gdopo.com/db"
	"gdopo.com/global"
	"gdopo.com/middleware"
	"gdopo.com/model"
	"gdopo.com/router"
	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRoute() *gin.Engine {
	global.Config = config.InitSysConfig() // 初始化配置文件

	global.MySQL = db.InitMySQL()                                // 链接数据库
	global.MySQL.Debug().AutoMigrate(model.AutoMigrateModels...) // 自动迁移数据库模型

	global.Casbin = middleware.InitCasbin() // 初始化 casbin

	app := gin.Default()                 // 初始化服务
	app.Use(gin.Logger())                // 日志服务
	app.Use(gin.Recovery())              // 日志服务
	app.NoRoute(global.NotFound)         // 处理 404
	v1 := app.Group("/api").Group("/v1") // 路由
	router.SwaggerRouter(v1)             // swagger api 文档
	router.AuthRouter(v1)                // 认证登录相关
	router.UserRouter(v1)                // 用户管理相关
	return app
}
