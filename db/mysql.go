// @File    :   mysql.go
// @Time    :   2020/11/24 22:33:44
// @Author  :   Lison
// @Version :   1.0
// @Contact :   linliangsuan@qq.com
// @License :   (C)Copyright 2017-2018, Liugroup-NLPR-CASIA
// @Desc    :   链接数据库

package db

import (
	"fmt"

	"gdopo.com/global"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitMySQL() *gorm.DB {
	conn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", global.Config.MySQL_User, global.Config.MySQL_Pass, global.Config.MySQL_Host, global.Config.MySQL_Port, global.Config.MySQL_DB)
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		fmt.Println(err)
	}
	db.SingularTable(true)
	return db

}
