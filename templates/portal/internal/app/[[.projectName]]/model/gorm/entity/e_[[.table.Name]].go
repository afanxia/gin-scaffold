[[set . "t_class" (.table.Name | singular | camel)]]
package entity

import (
	"context"

	"[[.project]]/internal/app/[[.projectName]]/schema"
	"[[.project]]/internal/app/common/model/gorm/entity"
	"[[.project]]/pkg/gormplus"
)

// Get[[.t_class]]DB 获取[[.table.Name]]存储
func Get[[.t_class]]DB(ctx context.Context, defDB *gormplus.DB) *gormplus.DB {
	return entity.GetDBWithModel(ctx, defDB, [[.t_class]]{})
}

// Schema[[.t_class]] [[.table.Name]]对象
type Schema[[.t_class]] schema.[[.t_class]]

// To[[.t_class]] 转换为[[.table.Name]]实体
func (a Schema[[.t_class]]) To[[.t_class]]() *[[.t_class]] {
	item := &[[.t_class]]{
		[[range .table.Columns]]
		[[- .Field | camel | lint]]:  a.[[.Field | camel | lint]],
		[[end]]
	}
	return item
}

// [[.t_class]] [[.table.Name]]实体
type [[.t_class]] struct {
	entity.Model
	[[range .table.Columns]]
	[[- .Field | camel | lint]]	[[convert "mysql" .Type (.Tag "gotype")]]	`db:"[[.Field]]"`
	[[end]]
}

func (a [[.t_class]]) String() string {
	return entity.ToString(a)
}

// TableName 表名
func (a [[.t_class]]) TableName() string {
	return a.Model.TableName("[[.table.Name]]")
}

// ToSchema[[.t_class]] 转换为[[.table.Name]]对象
func (a [[.t_class]]) ToSchema[[.t_class]]() *schema.[[.t_class]] {
	item := &schema.[[.t_class]]{
		[[range .table.Columns]]
		[[- .Field | camel | lint]]:  a.[[.Field | camel | lint]],
		[[end]]
	}
	return item
}

// [[.t_class]]s [[.table.Name]]列表
type [[.t_class]]s []*[[.t_class]]

// ToSchema[[.t_class]]s 转换为[[.table.Name]]对象列表
func (a [[.t_class]]s) ToSchema[[.t_class]]s() []*schema.[[.t_class]] {
	list := make([]*schema.[[.t_class]], len(a))
	for i, item := range a {
		list[i] = item.ToSchema[[.t_class]]()
	}
	return list
}
