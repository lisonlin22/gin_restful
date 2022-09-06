// @File    :   jwt.go
// @Time    :   2020/11/25 20:51:14
// @Author  :   Lison
// @Version :   1.0
// @Contact :   linliangsuan@qq.com
// @License :   (C)Copyright 2017-2018, Liugroup-NLPR-CASIA
// @Desc    :   Token 认证

package middleware

import (
	"errors"
	"strings"

	"gdopo.com/global"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var MySecret = []byte("Chinese baby")

// GenToken 生成JWT
func GenToken(MyClaims global.MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims) // 使用指定的签名方法创建签名对象
	return token.SignedString(MySecret)                          // 使用指定的secret签名并获得完整的编码后的字符串token
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*global.MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &global.MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*global.MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			global.Unauthorized(c)
			c.Abort()
			return
		}
		// 按空格分割
		token := strings.SplitN(authHeader, " ", 2)
		// token[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(token[0])
		if err != nil {
			global.Unauthorized(c)
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("user", mc)
		c.Next() // 后续的处理函数可以用过c.Get("user")来获取当前请求的用户信息
	}
}
