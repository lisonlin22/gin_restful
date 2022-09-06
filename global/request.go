// @File    :   request.go
// @Time    :   2020/11/24 22:07:03
// @Author  :   Lison
// @Version :   1.0
// @Contact :   linliangsuan@qq.com
// @License :   (C)Copyright 2017-2018, Liugroup-NLPR-CASIA
// @Desc    :   公共请求体
package global

import (
	"github.com/dgrijalva/jwt-go"
)

// 认证中间件请求体
type MyClaims struct {
	Username      string `json:"username" `
	AuthorityName string `json:"authorityName"`
	jwt.StandardClaims
}
