package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"InnoTaxi/internal/app/auth"
	"InnoTaxi/internal/app/errors"
)

func AuthMiddleware(forger auth.TokenForger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if err := forger.Decode(token); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidToken.Error())
		}
		ctx.Next()
	}
}
