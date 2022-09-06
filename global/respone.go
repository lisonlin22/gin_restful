// @File    :   respone.go
// @Time    :   2020/11/24 21:58:54
// @Author  :   Lison
// @Version :   1.0
// @Contact :   linliangsuan@qq.com
// @License :   (C)Copyright 2017-2018, Liugroup-NLPR-CASIA
// @Desc    :   数据接口返回函数

package global

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 返回成功
func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": []string{},
		"msg":  "",
	})
}
func SuccessData(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
		"msg":  "",
	})
}

// 返回失败
func Error(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": []string{},
		"msg":  "操作失败",
	})
}
func ErrorData(data interface{}, c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 1,
		"data": data,
		"msg":  "操作失败",
	})
}

// 404
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code": 404,
		"data": []string{},
		"msg":  "未找到！",
	})
}

// 401
func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code": 401,
		"data": []string{},
		"msg":  "认证未授权！",
	})
}

// 403
func NotForbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{
		"code": 403,
		"data": []string{},
		"msg":  "无权限访问！",
	})
}
