// @File    :   auth.go
// @Time    :   2020/11/25 21:12:37
// @Author  :   Lison
// @Version :   1.0
// @Contact :   linliangsuan@qq.com
// @License :   (C)Copyright 2017-2018, Liugroup-NLPR-CASIA
// @Desc    :   用户管理

package router

import (
	v1 "gdopo.com/api/v1"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	var view v1.UserController
	api := r.Group("/user")
	{
		api.GET("/", view.List)         // 用户列表
		api.POST("/", view.Create)      // 添加用户
		api.GET("/:Id", view.Retrieve)  // 查看用户
		api.PUT("/:Id", view.Put)       // 修改用户
		api.DELETE("/:Id", view.Delete) // 删除用户
	}
}
