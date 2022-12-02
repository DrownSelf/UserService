package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/DrownSelf/UserService/internal/auth"
	"github.com/DrownSelf/UserService/internal/handlers"
	"github.com/DrownSelf/UserService/internal/repositories"
)

type ApiV1 struct {
	handler   *handlers.Handler
	auth      auth.TokenAuth
	cacheRepo repositories.ICacheRepository
}

func NewApiV1(handler *handlers.Handler, tokenAuth auth.TokenAuth, repository repositories.ICacheRepository) *ApiV1 {
	return &ApiV1{handler: handler, auth: tokenAuth, cacheRepo: repository}
}

func (a *ApiV1) InitApiV1Groups(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	{
		orderGroup := v1.Group("/order")
		{
			orderGroup.POST("/rateUser", a.handler.RateUserFromOrder)
		}
		a.UserGroup(v1)
	}
}
