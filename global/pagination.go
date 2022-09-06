// @File    :   pagination.go
// @Time    :   2020/11/27 00:21:55
// @Author  :   Lison
// @Version :   1.0
// @Contact :   linliangsuan@qq.com
// @License :   (C)Copyright 2017-2018, Liugroup-NLPR-CASIA
// @Desc    :   分页器

package global

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// 分页器
func Pagination(c *gin.Context) (page int, size int, count int) {
	p := string(c.DefaultQuery("page", Config.Page))
	page, _ = strconv.Atoi(p)
	s := string(c.DefaultQuery("size", Config.Size))
	size, _ = strconv.Atoi(s)
	count = 0
	return (page - 1) * size, size, count
}
