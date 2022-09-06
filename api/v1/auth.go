// @File    :   auth.go
// @Time    :   2020/11/25 21:01:54
// @Author  :   Lison
// @Version :   1.0
// @Contact :   linliangsuan@qq.com
// @License :   (C)Copyright 2017-2018, Liugroup-NLPR-CASIA
// @Desc    :   认证登录

package v1

import (
	"net/http"
	"time"

	"gdopo.com/global"
	"gdopo.com/middleware"
	"gdopo.com/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 登录表单
type LoginForm struct {
	Username string `json:"username;required"` // 用户名 必须
	Password string `json:"password;required"` // 密  码 必须
}

// @Tags 认证登录
// @Summary 登录
// @accept application/json
// @Produce application/json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 {string} string "{"code": 0, "data":{"token": "", },"msg":""}"
// @Router /api/v1/login/ [post]
func Login(c *gin.Context) {
	var params LoginForm
	err := c.ShouldBind(&params)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "无效的参数",
			"data": []string{},
		})
		return
	}
	// md5 加密比对
	params.Password = global.Md5Hash(params.Password)
	// 校验用户名和密码是否正确
	var user model.Users
	if err := global.MySQL.Preloads("Authority").Where("username =? and password = ? and active = 1", params.Username, params.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "登录失败",
			"data": []string{},
		})
		return
	}
	var authority model.Authority
	global.MySQL.Model(&authority).Where("authority_id = ?", user.AuthorityId).First(&authority)
	Claims := global.MyClaims{
		Username:      user.Username,
		AuthorityName: authority.AuthorityName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "linliangsuan@qq.com",           //签名的发行者
		},
	}
	// 生成 Token
	tokenString, _ := middleware.GenToken(Claims)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登录成功",
		"data": gin.H{"token": tokenString},
	})
	return
}

// @Tags 认证登录
// @Summary 注销
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"code": 0, "data":{"token": "", },"msg":""}"
// @Router /api/v1/logout/ [post]
func Logout(c *gin.Context) {
	// 获取绘画信息
	user, _ := c.Get("user")
	MyClaims := user.(*global.MyClaims)
	Claims := global.MyClaims{
		Username: MyClaims.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 1),    // 过期时间 一小时
			Issuer:    "linliangsuan@qq.com",           //签名的发行者
		},
	}
	// 重新生成token
	tokenString, _ := middleware.GenToken(Claims)
	if len(tokenString) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "注销失败",
			"data": []string{},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "注销成功！",
		"data": []string{},
	})
	return

}
