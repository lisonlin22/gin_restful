// @File    :   model.go
// @Time    :   2020/11/24 22:57:25
// @Author  :   Lison
// @Version :   1.0
// @Contact :   linliangsuan@qq.com
// @License :   (C)Copyright 2017-2018, Liugroup-NLPR-CASIA
// @Desc    :   数据库模型

package model

import (
	"gdopo.com/global"
)

// 用户列表
type Users struct {
	global.Model
	Username    string    `form:"username" json:"username" gorm:"comment:'用户登录名'"`
	Password    string    `json:"-"  gorm:"comment:'用户登录密码'"`
	NickName    string    `form:"nickname" json:"nickname" gorm:"default:'系统用户';comment:'用户昵称'" `
	Active      bool      `form:"active" json:"active" gorm:"default:true;comment:'激活状态'" `
	HeaderImg   string    `form:"headerImg" json:"headerImg" gorm:"default:'http://qmplusimg.henrongyi.top/head.png';comment:'用户头像'"`
	Authority   Authority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:'用户角色'"`
	AuthorityId string    `form:"authorityId" json:"authorityId" gorm:"default:0;comment:'用户角色ID'"`
}

// 用户角色
type Authority struct {
	global.Model
	AuthorityId   string `json:"authorityId" gorm:"not null;unique;primary_key;comment:'角色ID';size:90"`
	AuthorityName string `json:"authorityName" gorm:"comment:'角色名'"`
	Menus         []Menu `json:"menus" gorm:"many2many:sys_authority_menus;"`
}

// 系统菜单
type Menu struct {
	global.Model
	MenuLevel  uint   `json:"-"`
	Path       string `json:"path" gorm:"comment:'路由path'"`
	Name       string `json:"name" gorm:"comment:'路由name'"`
	Hidden     bool   `json:"hidden" gorm:"comment:'是否在列表隐藏'"`
	Component  string `json:"component" gorm:"comment:'对应前端文件路径'"`
	Sort       int    `json:"sort" gorm:"comment:'排序标记'"`
	Meta       `json:"meta" gorm:"comment:'附加属性'"`
	Authoritys []Authority `json:"authoritys" gorm:"many2many:sys_authority_menus;"`
	Children   []Menu      `json:"children" gorm:"-"`
}

// 菜单附加元素
type Meta struct {
	KeepAlive bool   `json:"keepAlive" gorm:"comment:'是否缓存'"`
	Title     string `json:"title" gorm:"comment:'菜单名'"`
	Icon      string `json:"icon" gorm:"comment:'菜单图标'"`
}
