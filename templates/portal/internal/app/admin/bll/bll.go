package bll

import (
	"context"
	"sync"

	"[[.project]]/internal/app/admin/schema"
	"[[.project]]/internal/app/config"
	"[[.project]]/internal/app/common/model"
	icontext "[[.project]]/internal/app/context"
	"[[.project]]/pkg/util"
)

// GetUserID 获取用户ID
func GetUserID(ctx context.Context) string {
	userID, ok := icontext.FromUserID(ctx)
	if ok {
		return userID
	}
	return ""
}

// TransFunc 定义事务执行函数
type TransFunc func(context.Context) error

// ExecTrans 执行事务
func ExecTrans(ctx context.Context, transModel model.ITrans, fn TransFunc) error {
	if _, ok := icontext.FromTrans(ctx); ok {
		return fn(ctx)
	}
	trans, err := transModel.Begin(ctx)
	if err != nil {
		return err
	}

	err = fn(icontext.NewTrans(ctx, trans))
	if err != nil {
		_ = transModel.Rollback(ctx, trans)
		return err
	}
	return transModel.Commit(ctx, trans)
}

var (
	rootUser     *schema.User
	rootUserOnce sync.Once
)

// GetRootUser 获取root用户
func GetRootUser() *schema.User {
	rootUserOnce.Do(func() {
		user := config.GetGlobalConfig().Root
		rootUser = &schema.User{
			RecordID: user.UserName,
			UserName: user.UserName,
			RealName: user.RealName,
			Password: util.MD5HashString(user.Password),
		}
	})
	return rootUser
}

// CheckIsRootUser 检查是否是root用户
func CheckIsRootUser(ctx context.Context, recordID string) bool {
	return GetRootUser().RecordID == recordID
}

