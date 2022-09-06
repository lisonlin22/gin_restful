// @File    :   auth.go
// @Time    :   2020/11/25 21:12:37
// @Author  :   Lison
// @Version :   1.0
// @Contact :   linliangsuan@qq.com
// @License :   (C)Copyright 2017-2018, Liugroup-NLPR-CASIA
// @Desc    :   登录认证

package router

import (
	v1 "gdopo.com/api/v1"
	"gdopo.com/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup) {
	// 登录
	api := r.Group("/login")
	{
		api.POST("/", v1.Login)
	}
	// 注销
	logout := r.Group("/logout").Use(middleware.JWTAuthMiddleware())
	{
		logout.POST("/", v1.Logout)
	}
}
