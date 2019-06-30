[[set . "t_class" (.table.Name | singular | camel)]]
package bll

import (
	"context"

	"[[.project]]/internal/app/[[.projectName]]/model"
	"[[.project]]/internal/app/admin/bll"
	"[[.project]]/internal/app/[[.projectName]]/schema"
	cmodel "[[.project]]/internal/app/common/model"
	cschema "[[.project]]/internal/app/common/schema"
	"[[.project]]/pkg/errors"
	"[[.project]]/pkg/util"
)

// New[[.t_class]] 创建[[.table.Name]]管理实例
func New[[.t_class]](m *cmodel.Common) *[[.t_class]] {
	return &[[.t_class]]{
		[[.t_class]]Model: m.[[.t_class]],
	}
}

// [[.t_class]] [[.table.Name]]管理
type [[.t_class]] struct {
	[[.t_class]]Model model.I[[.t_class]]
}

// QueryPage 查询分页数据
func (a *[[.t_class]]) QueryPage(ctx context.Context, params schema.[[.t_class]]QueryParam, pp *cschema.PaginationParam) ([]*schema.[[.t_class]], *cschema.PaginationResult, error) {
	result, err := a.[[.t_class]]Model.Query(ctx, params, schema.[[.t_class]]QueryOptions{PageParam: pp})
	if err != nil {
		return nil, nil, err
	}
	return result.Data, result.PageResult, nil
}

// Get 查询指定数据
func (a *[[.t_class]]) Get(ctx context.Context, recordID string) (*schema.[[.t_class]], error) {
	item, err := a.[[.t_class]]Model.Get(ctx, recordID)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *[[.t_class]]) checkCode(ctx context.Context, code string) error {
	result, err := a.[[.t_class]]Model.Query(ctx, schema.[[.t_class]]QueryParam{
		Code: code,
	}, schema.[[.t_class]]QueryOptions{
		PageParam: &cschema.PaginationParam{PageSize: -1},
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.NewBadRequestError("编号已经存在")
	}
	return nil
}

// Create 创建数据
func (a *[[.t_class]]) Create(ctx context.Context, item schema.[[.t_class]]) (*schema.[[.t_class]], error) {
	err := a.checkCode(ctx, item.Code)
	if err != nil {
		return nil, err
	}

	item.RecordID = util.MustUUID()
	item.Creator = bll.GetUserID(ctx)
	err = a.[[.t_class]]Model.Create(ctx, item)
	if err != nil {
		return nil, err
	}
	return a.Get(ctx, item.RecordID)
}

// Update 更新数据
func (a *[[.t_class]]) Update(ctx context.Context, recordID string, item schema.[[.t_class]]) (*schema.[[.t_class]], error) {
	oldItem, err := a.[[.t_class]]Model.Get(ctx, recordID)
	if err != nil {
		return nil, err
	} else if oldItem == nil {
		return nil, errors.ErrNotFound
	} else if oldItem.Code != item.Code {
		err := a.checkCode(ctx, item.Code)
		if err != nil {
			return nil, err
		}
	}

	err = a.[[.t_class]]Model.Update(ctx, recordID, item)
	if err != nil {
		return nil, err
	}
	return a.Get(ctx, recordID)
}

// Delete 删除数据
func (a *[[.t_class]]) Delete(ctx context.Context, recordID string) error {
	err := a.[[.t_class]]Model.Delete(ctx, recordID)
	if err != nil {
		return err
	}
	return nil
}

// UpdateStatus 更新状态
func (a *[[.t_class]]) UpdateStatus(ctx context.Context, recordID string, status int) error {
	return a.[[.t_class]]Model.UpdateStatus(ctx, recordID, status)
}
