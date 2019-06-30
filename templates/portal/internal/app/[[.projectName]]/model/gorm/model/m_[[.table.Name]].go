[[set . "t_class" (.table.Name | singular | camel)]]
package model

import (
	"context"
	"fmt"

	"[[.project]]/internal/app/[[.projectName]]/model/gorm/entity"
	"[[.project]]/internal/app/[[.projectName]]/schema"
	"[[.project]]/internal/app/common/model/gorm/model"
	"[[.project]]/pkg/errors"
	"[[.project]]/pkg/gormplus"
	"[[.project]]/pkg/logger"
)

// New[[.t_class]] 创建[[.table.Name]]存储实例
func New[[.t_class]](db *gormplus.DB) *[[.t_class]] {
	return &[[.t_class]]{db}
}

// [[.t_class]] [[.table.Name]]存储
type [[.t_class]] struct {
	db *gormplus.DB
}

func (a *[[.t_class]]) getFuncName(name string) string {
	return fmt.Sprintf("gorm.model.[[.t_class]].%s", name)
}

func (a *[[.t_class]]) getQueryOption(opts ...schema.[[.t_class]]QueryOptions) schema.[[.t_class]]QueryOptions {
	var opt schema.[[.t_class]]QueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *[[.t_class]]) Query(ctx context.Context, params schema.[[.t_class]]QueryParam, opts ...schema.[[.t_class]]QueryOptions) (*schema.[[.t_class]]QueryResult, error) {
	span := logger.StartSpan(ctx, "查询数据", a.getFuncName("Query"))
	defer span.Finish()

	db := entity.Get[[.t_class]]DB(ctx, a.db).DB
	if v := params.Code; v != "" {
		db = db.Where("code=?", v)
	}
	if v := params.LikeCode; v != "" {
		db = db.Where("code LIKE ?", "%"+v+"%")
	}
	if v := params.LikeName; v != "" {
		db = db.Where("name LIKE ?", "%"+v+"%")
	}
	if v := params.Status; v > 0 {
		db = db.Where("status=?", v)
	}
	db = db.Order("id DESC")

	opt := a.getQueryOption(opts...)
	var list entity.[[.t_class]]s
	pr, err := model.WrapPageQuery(db, opt.PageParam, &list)
	if err != nil {
		span.Errorf(err.Error())
		return nil, errors.New("查询数据发生错误")
	}
	qr := &schema.[[.t_class]]QueryResult{
		PageResult: pr,
		Data:       list.ToSchema[[.t_class]]s(),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *[[.t_class]]) Get(ctx context.Context, recordID string, opts ...schema.[[.t_class]]QueryOptions) (*schema.[[.t_class]], error) {
	span := logger.StartSpan(ctx, "查询指定数据", a.getFuncName("Get"))
	defer span.Finish()

	db := entity.Get[[.t_class]]DB(ctx, a.db).Where("record_id=?", recordID)
	var item entity.[[.t_class]]
	ok, err := a.db.FindOne(db, &item)
	if err != nil {
		span.Errorf(err.Error())
		return nil, errors.New("查询指定数据发生错误")
	} else if !ok {
		return nil, nil
	}

	return item.ToSchema[[.t_class]](), nil
}

// Create 创建数据
func (a *[[.t_class]]) Create(ctx context.Context, item schema.[[.t_class]]) error {
	span := logger.StartSpan(ctx, "创建数据", a.getFuncName("Create"))
	defer span.Finish()

	[[.table.Name]] := entity.Schema[[.t_class]](item).To[[.t_class]]()
	result := entity.Get[[.t_class]]DB(ctx, a.db).Create([[.table.Name]])
	if err := result.Error; err != nil {
		span.Errorf(err.Error())
		return errors.New("创建数据发生错误")
	}
	return nil
}

// Update 更新数据
func (a *[[.t_class]]) Update(ctx context.Context, recordID string, item schema.[[.t_class]]) error {
	span := logger.StartSpan(ctx, "更新数据", a.getFuncName("Update"))
	defer span.Finish()

	[[.table.Name]] := entity.Schema[[.t_class]](item).To[[.t_class]]()
	result := entity.Get[[.t_class]]DB(ctx, a.db).Where("record_id=?", recordID).Omit("record_id", "creator").Updates([[.table.Name]])
	if err := result.Error; err != nil {
		span.Errorf(err.Error())
		return errors.New("更新数据发生错误")
	}
	return nil
}

// Delete 删除数据
func (a *[[.t_class]]) Delete(ctx context.Context, recordID string) error {
	span := logger.StartSpan(ctx, "删除数据", a.getFuncName("Delete"))
	defer span.Finish()

	result := entity.Get[[.t_class]]DB(ctx, a.db).Where("record_id=?", recordID).Delete(entity.[[.t_class]]{})
	if err := result.Error; err != nil {
		span.Errorf(err.Error())
		return errors.New("删除数据发生错误")
	}
	return nil
}

// UpdateStatus 更新状态
func (a *[[.t_class]]) UpdateStatus(ctx context.Context, recordID string, status int) error {
	span := logger.StartSpan(ctx, "更新状态", a.getFuncName("UpdateStatus"))
	defer span.Finish()

	result := entity.Get[[.t_class]]DB(ctx, a.db).Where("record_id=?", recordID).Update("status", status)
	if err := result.Error; err != nil {
		span.Errorf(err.Error())
		return errors.New("更新状态发生错误")
	}
	return nil
}
