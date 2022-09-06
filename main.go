// @File    :   main.go
// @Time    :   2020/11/23 22:51:23
// @Author  :   Lison
// @Version :   1.0
// @Contact :   linliangsuan@qq.com
// @License :   (C)Copyright 2017-2018, Liugroup-NLPR-CASIA
// @Desc    :   None

package main

import (
	"fmt"

	"gdopo.com/app"
	"gdopo.com/global"
)

func main() {
	app := app.InitRoute()
	app.Run(fmt.Sprintf("%s:%s", global.Config.Listen, global.Config.Port))

}
