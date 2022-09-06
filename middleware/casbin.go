// @File    :   casbin.go
// @Time    :   2020/11/25 21:35:40
// @Author  :   Lison
// @Version :   1.0
// @Contact :   linliangsuan@qq.com
// @License :   (C)Copyright 2017-2018, Liugroup-NLPR-CASIA
// @Desc    :   权限控制

package middleware

import (
	"gdopo.com/global"
	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
	"github.com/gin-gonic/gin"
)

// 初始化 casbin 权限
func InitCasbin() *casbin.Enforcer {
	// mysql 适配器
	adapter := gormadapter.NewAdapterByDB(global.MySQL)
	// 通过mysql适配器新建一个enforcer
	Enforcer := casbin.NewEnforcer("D:/gitee/gin_restful/config/rbac_model.conf", adapter)
	// 日志记录
	Enforcer.EnableLog(true)
	return Enforcer
}

// Casbin 权限策略
func CasbinPerms() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")
		sub := user.(*global.MyClaims) // 对象
		obj := c.Request.URL.Path      // 访问URL
		act := c.Request.Method        // 操作
		result, err := global.Casbin.EnforceSafe(sub.Username, obj, act)

		if err != nil || !result { // 判断是否有权限
			global.NotForbidden(c)
			c.Abort()
			return
		}
		c.Next()
	}
}
