package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/DrownSelf/UserService/internal/middlewares"
)

func (a *ApiV1) UserGroup(group *gin.RouterGroup) {
	userGroup := group.Group("/user")
	{
		userGroup.POST("/register", a.handler.Register)
		userGroup.POST("/login", a.handler.LogIn)
		authGroup := userGroup.Group("", middlewares.TokenDecoderMiddleware(a.auth, a.cacheRepo))
		{
			authGroup.POST("/orderTaxi", a.handler.OrderTaxi)
			authGroup.POST("/rateOrder", a.handler.RateRide)
			authGroup.GET("/logout", a.handler.LogOut)
			authGroup.GET("", a.handler.GetUser)
			authGroup.PUT("", a.handler.UpdateUser)
			authGroup.DELETE("", a.handler.DeleteUser)
		}
	}
}
