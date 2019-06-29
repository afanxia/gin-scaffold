package model

import (
	adminmodel "[[.project]]/internal/app/admin/model"
)

// Common 提供统一的存储接口
type Common struct {
	Trans ITrans
	Demo  adminmodel.IDemo
	Menu  adminmodel.IMenu
	Role  adminmodel.IRole
	User  adminmodel.IUser
}
