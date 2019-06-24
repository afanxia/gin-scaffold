/*
Package app [[.projectName]] 生成swagger文档

文档规则请参考：https://github.com/teambition/swaggo/wiki/Declarative-Comments-Format

使用方式：

	go get -u -v github.com/teambition/swaggo
	swaggo -s ./internal/app/swagger.go -p . -o ./internal/app/swagger
*/
package app

import (
	// API控制器
	_ "[[.project]]/internal/app/admin/routers/api/ctl"
)

// @Version 0.0.1
// @Title GinScaffold
// @Description gin scaffold
// @Schemes http,https
// @Host 127.0.0.1:10088
// @BasePath /
// @Name Afanxia
// @Contact afanxia@163.com
// @Consumes json
// @Produces json
