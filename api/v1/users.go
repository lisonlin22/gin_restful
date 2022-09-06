// @File    :   users.go
// @Time    :   2020/11/26 21:22:36
// @Author  :   Lison
// @Version :   1.0
// @Contact :   linliangsuan@qq.com
// @License :   (C)Copyright 2017-2018, Liugroup-NLPR-CASIA
// @Desc    :   用户管理

package v1

import (
	"net/http"

	"gdopo.com/global"
	"gdopo.com/model"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Queryset model.Users   // 模型序列化
	Params   model.Users   // 读取参数用于过滤器
	Result   model.Users   // 返回单条数据
	Results  []model.Users // 返回多条数据
}

// @Tags 用户管理
// @Summary 用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users false "用户名, 激活状态，头像，角色"
// @Success 200 {string} string "{"code": 0, "data":[],"msg":""}"
// @Router /api/v1/user/ [get]
func (t *UserController) List(c *gin.Context) { // 过滤器
	_ = c.ShouldBindQuery(&t.Params)          // 读取参数
	page, size, count := global.Pagination(c) // 分页器
	if err := global.MySQL.Preload("Authority").Model(&t.Queryset).Where(&t.Params).Count(&count).Limit(size).Offset(page).Find(&t.Results).Error; err != nil {
		global.Error(c)
		return
	}
	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"data":  t.Results,
		"msg":   "",
		"count": count,
	})
	return
}

// @Tags 用户管理
// @Summary 添加用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users false "用户名, 激活状态，头像，角色"
// @Success 200 {string} string "{"code": 0, "data":{},"msg":""}"
// @Router /api/v1/user/ [post]
func (t *UserController) Create(c *gin.Context) {
	_ = c.ShouldBindJSON(&t.Params) // 读取参数
	if err := global.MySQL.Model(&t.Queryset).Create(&t.Params).Error; err != nil {
		global.Error(c)
		return
	}
	// 返回结果
	global.SuccessData(&t.Params, c)
	return
}

// @Tags 用户管理
// @Summary 查看用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"code": 0, "data":{},"msg":""}"
// @Router /api/v1/user/:id/ [get]
func (t *UserController) Retrieve(c *gin.Context) {
	Id := c.Param("Id") // 读取参数
	if err := global.MySQL.Model(&t.Queryset).First(&t.Result, Id).Error; err != nil {
		global.NotFound(c)
		return
	}
	// 返回结果
	global.SuccessData(&t.Result, c)
	return
}

// @Tags 用户管理
// @Summary 修改用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users false "用户名, 激活状态，头像，角色"
// @Success 200 {string} string "{"code": 0, "data":{},"msg":""}"
// @Router /api/v1/user/:id/ [put]
func (t *UserController) Put(c *gin.Context) {
	Id := c.Param("Id")             // 读取主键
	_ = c.ShouldBindJSON(&t.Params) // 读取参数
	if err := global.MySQL.Model(&t.Queryset).First(&t.Result, Id).Update(&t.Params).Error; err != nil {
		global.Error(c)
		return
	}
	// 返回结果
	global.SuccessData(&t.Params, c)
	return
}

// @Tags 用户管理
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users false "用户名, 激活状态，头像，角色"
// @Success 200 {string} string "{"code": 0, "data":{},"msg":""}"
// @Router /api/v1/user/:id/ [delete]
func (t *UserController) Delete(c *gin.Context) {
	Id := c.Param("Id") // 读取主键
	if err := global.MySQL.Model(&t.Queryset).First(&t.Result, Id).Delete(&t.Result).Error; err != nil {
		global.Error(c)
		return
	}
	// 返回结果
	global.Success(c)
	return
}
