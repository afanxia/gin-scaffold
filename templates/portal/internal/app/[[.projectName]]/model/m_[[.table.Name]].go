[[set . "t_class" (.table.Name | singular | camel)]]
package model

import (
	"context"

	"[[.project]]/internal/app/[[.projectName]]/schema"
)

// I[[.t_class]] [[.table.Name]]存储接口
type I[[.t_class]] interface {
	// 查询数据
	Query(ctx context.Context, params schema.[[.t_class]]QueryParam, opts ...schema.[[.t_class]]QueryOptions) (*schema.[[.t_class]]QueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, recordID string, opts ...schema.[[.t_class]]QueryOptions) (*schema.[[.t_class]], error)
	// 创建数据
	Create(ctx context.Context, item schema.[[.t_class]]) error
	// 更新数据
	Update(ctx context.Context, recordID string, item schema.[[.t_class]]) error
	// 删除数据
	Delete(ctx context.Context, recordID string) error
	// 更新状态
	UpdateStatus(ctx context.Context, recordID string, status int) error
}
