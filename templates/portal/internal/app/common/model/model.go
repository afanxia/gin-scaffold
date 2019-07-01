[[set . "t_class" (.table.Name | singular | camel)]]
package model

import (
	adminmodel "[[.project]]/internal/app/admin/model"
	[[.projectName]]model "[[.project]]/internal/app/[[.projectName]]/model"
)

// Common 提供统一的存储接口
type Common struct {
	Trans ITrans
	Demo  adminmodel.IDemo
	Menu  adminmodel.IMenu
	Role  adminmodel.IRole
	User  adminmodel.IUser
	[[range $t := .tables]]
	[[- $t.Name | singular | camel]] [[$.projectName]]model.I[[$t.Name | singular | camel]]
	[[end]]
}
