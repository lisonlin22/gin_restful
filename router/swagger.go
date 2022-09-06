// @File    :   swagger.go
// @Time    :   2020/11/24 22:50:58
// @Author  :   Lison
// @Version :   1.0
// @Contact :   linliangsuan@qq.com
// @License :   (C)Copyright 2017-2018, Liugroup-NLPR-CASIA
// @Desc    :   Swagger API 文档路由

package router

import (
	_ "gdopo.com/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerRouter(r *gin.RouterGroup) {
	api := r.Group("/swagger")
	{
		api.GET("/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))
	}
}
