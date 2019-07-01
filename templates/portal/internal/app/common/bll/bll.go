[[set . "t_class" (.table.Name | singular | camel)]]
package bll

import (
	adminbll "[[.project]]/internal/app/admin/bll"
	[[.projectName]]bll "[[.project]]/internal/app/[[.projectName]]/bll"
	"[[.project]]/internal/app/common/model"
	"[[.project]]/pkg/auth"
	"github.com/casbin/casbin"
)

// Common 提供统一的业务逻辑处理
type Common struct {
	Demo  *adminbll.Demo
	Login *adminbll.Login
	Menu  *adminbll.Menu
	Role  *adminbll.Role
	User  *adminbll.User
	[[range $t := .tables]]
	[[- $t.Name | singular | camel]] *[[$.projectName]]bll.[[$t.Name | singular | camel]]
	[[end]]
}

// NewCommon 创建统一的业务逻辑处理
func NewCommon(m *model.Common, a auth.Auther, e *casbin.Enforcer) *Common {
	return &Common{
		Demo:  adminbll.NewDemo(m),
		Login: adminbll.NewLogin(m, a),
		Menu:  adminbll.NewMenu(m),
		Role:  adminbll.NewRole(m, e),
		User:  adminbll.NewUser(m, e),
		[[range $t := .tables]]
		[[- $t.Name | singular | camel]]: [[$.projectName]]bll.New[[$t.Name | singular | camel]](m),
		[[end]]
	}
}
