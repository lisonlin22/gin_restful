// @File    :   config.go
// @Time    :   2020/11/24 21:06:48
// @Author  :   Lison
// @Version :   1.0
// @Contact :   linliangsuan@qq.com
// @License :   (C)Copyright 2017-2018, Liugroup-NLPR-CASIA
// @Desc    :   项目配置文件

package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type SysConfig struct {
	Listen     string // 监听地址
	Port       string // 监听端口
	MySQL_Host string
	MySQL_Port string
	MySQL_User string
	MySQL_Pass string
	MySQL_DB   string
	Redis_Host string
	Redis_Port string
	Redis_Auth string
	Redis_DB   string
	Page       string
	Size       string
}

// 读取系统配置文件
func InitSysConfig() SysConfig {
	cfg, err := ini.Load("D:/gitee/gin_restful/config/config.ini")
	if err != nil {
		fmt.Println("error")
	}
	var s SysConfig
	s.Port = cfg.Section("server").Key("port").String()
	s.Listen = cfg.Section("server").Key("listen").String()
	s.MySQL_Host = cfg.Section("mysql").Key("host").String()
	s.MySQL_Port = cfg.Section("mysql").Key("port").String()
	s.MySQL_User = cfg.Section("mysql").Key("username").String()
	s.MySQL_Pass = cfg.Section("mysql").Key("password").String()
	s.MySQL_DB = cfg.Section("mysql").Key("database").String()
	s.Redis_Host = cfg.Section("redis").Key("host").String()
	s.Redis_Port = cfg.Section("redis").Key("port").String()
	s.Redis_Auth = cfg.Section("redis").Key("auth").String()
	s.Redis_DB = cfg.Section("redis").Key("database").String()
	s.Page = cfg.Section("pagination").Key("page").String()
	s.Size = cfg.Section("pagination").Key("size").String()
	return s
}
