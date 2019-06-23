package middleware

import (
	"github.com/afanxia/gin-scaffold/internal/app/config"
	"github.com/afanxia/gin-scaffold/internal/app/ginplus"
	"github.com/afanxia/gin-scaffold/pkg/auth"
	"github.com/afanxia/gin-scaffold/pkg/errors"
	"github.com/afanxia/gin-scaffold/pkg/logger"
	"github.com/gin-gonic/gin"
)

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(a auth.Auther, skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userID string
		if t := ginplus.GetToken(c); t != "" {
			id, err := a.ParseUserID(t)
			if err != nil {
				if err == auth.ErrInvalidToken {
					ginplus.ResError(c, errors.NewUnauthorizedError())
					return
				}
				logger.StartSpan(ginplus.NewContext(c), "用户授权中间件", "UserAuthMiddleware").Errorf(err.Error())
				ginplus.ResError(c, errors.NewInternalServerError())
				return
			}
			userID = id
		}

		if userID != "" {
			c.Set(ginplus.UserIDKey, userID)
		}

		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}

		if userID == "" {
			if config.GetGlobalConfig().RunMode == "debug" {
				c.Set(ginplus.UserIDKey, config.GetGlobalConfig().Root.UserName)
				c.Next()
				return
			}
			ginplus.ResError(c, errors.NewUnauthorizedError("用户未登录"))
		}
	}
}
