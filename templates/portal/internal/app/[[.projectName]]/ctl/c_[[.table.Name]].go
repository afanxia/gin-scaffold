[[set . "t_class" (.table.Name | singular | camel)]]
[[set . "t_plural" (.table.Name | plural)]]
package ctl

import (
	"[[.project]]/internal/app/[[.projectName]]/bll"
	cbll "[[.project]]/internal/app/common/bll"
	"[[.project]]/internal/app/[[.projectName]]/schema"
	"[[.project]]/internal/app/ginplus"
	"[[.project]]/pkg/errors"
	"[[.project]]/pkg/util"
	"github.com/gin-gonic/gin"
)

// New[[.t_class]] 创建[[.table.Name]]控制器
func New[[.t_class]](b *cbll.Common) *[[.t_class]] {
	return &[[.t_class]]{
		[[.t_class]]Bll: b.[[.t_class]],
	}
}

// [[.t_class]] [[.table.Name]]
// @Name [[.t_class]]
// @Description [[.table.Name]]
type [[.t_class]] struct {
	[[.t_class]]Bll *bll.[[.t_class]]
}

// Query 查询数据
func (a *[[.t_class]]) Query(c *gin.Context) {
	switch c.Query("q") {
	case "page":
		a.QueryPage(c)
	default:
		ginplus.ResError(c, errors.NewBadRequestError("未知的查询类型"))
	}
}

// QueryPage 查询分页数据
// @Summary 查询分页数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param current query int true "分页索引" 1
// @Param pageSize query int true "分页大小" 10
// @Param code query string false "编号"
// @Param name query string false "名称"
// @Param status query int false "状态(1:启用 2:停用)"
// @Success 200 []schema.[[.t_class]] "查询结果：{list:列表数据,pagination:{current:页索引,pageSize:页大小,total:总数量}}"
// @Failure 400 schema.HTTPError "{error:{code:0,message:未知的查询类型}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router GET /api/v1/[[.t_plural]]?q=page
func (a *[[.t_class]]) QueryPage(c *gin.Context) {
	var params schema.[[.t_class]]QueryParam
	params.LikeCode = c.Query("code")
	params.LikeName = c.Query("name")
	params.Status = util.S(c.Query("status")).Int()

	items, pr, err := a.[[.t_class]]Bll.QueryPage(ginplus.NewContext(c), params, ginplus.GetPaginationParam(c))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	ginplus.ResPage(c, items, pr)
}

// Get 查询指定数据
// @Summary 查询指定数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Success 200 schema.[[.t_class]]
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 404 schema.HTTPError "{error:{code:0,message:资源不存在}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router GET /api/v1/[[.t_plural]]/{id}
func (a *[[.t_class]]) Get(c *gin.Context) {
	item, err := a.[[.t_class]]Bll.Get(ginplus.NewContext(c), c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, item)
}

// Create 创建数据
// @Summary 创建数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param body body schema.[[.t_class]] true
// @Success 200 schema.[[.t_class]]
// @Failure 400 schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router POST /api/v1/[[.t_plural]]
func (a *[[.t_class]]) Create(c *gin.Context) {
	var item schema.[[.t_class]]
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	nitem, err := a.[[.t_class]]Bll.Create(ginplus.NewContext(c), item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, nitem)
}

// Update 更新数据
// @Summary 更新数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Param body body schema.[[.t_class]] true
// @Success 200 schema.[[.t_class]]
// @Failure 400 schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router PUT /api/v1/[[.t_plural]]/{id}
func (a *[[.t_class]]) Update(c *gin.Context) {
	var item schema.[[.t_class]]
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	nitem, err := a.[[.t_class]]Bll.Update(ginplus.NewContext(c), c.Param("id"), item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, nitem)
}

// Delete 删除数据
// @Summary 删除数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Success 200 schema.HTTPStatus "{status:OK}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router DELETE /api/v1/[[.t_plural]]/{id}
func (a *[[.t_class]]) Delete(c *gin.Context) {
	err := a.[[.t_class]]Bll.Delete(ginplus.NewContext(c), c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

// Enable 启用数据
// @Summary 启用数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Success 200 schema.HTTPStatus "{status:OK}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router PATCH /api/v1/[[.t_plural]]/{id}/enable
func (a *[[.t_class]]) Enable(c *gin.Context) {
	err := a.[[.t_class]]Bll.UpdateStatus(ginplus.NewContext(c), c.Param("id"), 1)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

// Disable 禁用数据
// @Summary 禁用数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Success 200 schema.HTTPStatus "{status:OK}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router PATCH /api/v1/[[.t_plural]]/{id}/disable
func (a *[[.t_class]]) Disable(c *gin.Context) {
	err := a.[[.t_class]]Bll.UpdateStatus(ginplus.NewContext(c), c.Param("id"), 2)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}
