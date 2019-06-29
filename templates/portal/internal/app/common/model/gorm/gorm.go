package gorm

import (
	adminentity "[[.project]]/internal/app/admin/model/gorm/entity"
	"[[.project]]/internal/app/common/model"
	"[[.project]]/internal/app/common/model/gorm/entity"
	adminmodel "[[.project]]/internal/app/admin/model/gorm/model"
	cmodel "[[.project]]/internal/app/common/model/gorm/model"
	"[[.project]]/pkg/gormplus"
)

// SetTablePrefix 设定表名前缀
func SetTablePrefix(prefix string) {
	entity.SetTablePrefix(prefix)
}

// AutoMigrate 自动映射数据表
func AutoMigrate(db *gormplus.DB) error {
	return db.AutoMigrate(
		new(adminentity.User),
		new(adminentity.UserRole),
		new(adminentity.Role),
		new(adminentity.RoleMenu),
		new(adminentity.Menu),
		new(adminentity.MenuAction),
		new(adminentity.MenuResource),
		new(adminentity.Demo),
	).Error
}

// NewModel 创建gorm存储，实现统一的存储接口
func NewModel(db *gormplus.DB) *model.Common {
	return &model.Common{
		Trans: cmodel.NewTrans(db),
		Demo:  adminmodel.NewDemo(db),
		Menu:  adminmodel.NewMenu(db),
		Role:  adminmodel.NewRole(db),
		User:  adminmodel.NewUser(db),
	}
}
