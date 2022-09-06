// @File    :   global.go
// @Time    :   2020/11/24 21:07:58
// @Author  :   Lison
// @Version :   1.0
// @Contact :   linliangsuan@qq.com
// @License :   (C)Copyright 2017-2018, Liugroup-NLPR-CASIA
// @Desc    :   全局变量

package global

import (
	"crypto/md5"
	"encoding/hex"

	"gdopo.com/config"
	"github.com/casbin/casbin"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

var (
	Config config.SysConfig
	MySQL  *gorm.DB
	Casbin *casbin.Enforcer
	Redis  *redis.Pool
)

// md5 加密
func Md5Hash(String string) (hash string) {
	h := md5.New()
	h.Write([]byte(String)) // 需要加密的字符串
	cipherStr := h.Sum(nil)
	hash = hex.EncodeToString(cipherStr)
	return
}
