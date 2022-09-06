// @File    :   migrate.go
// @Time    :   2020/11/24 22:47:01
// @Author  :   Lison
// @Version :   1.0
// @Contact :   linliangsuan@qq.com
// @License :   (C)Copyright 2017-2018, Liugroup-NLPR-CASIA
// @Desc    :   自动迁移数据库模型

package model

var AutoMigrateModels = []interface{}{
	&Users{},
	&Authority{},
	&Menu{},
	&Meta{},
}
