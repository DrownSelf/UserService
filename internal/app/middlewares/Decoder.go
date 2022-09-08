package middlewares

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"InnoTaxi/internal/app/appErrors"
	"InnoTaxi/internal/app/auth"
	"InnoTaxi/internal/app/repositories"
)

func TokenDecoderMiddleware(forger auth.TokenForger, repository repositories.ICacheRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if repository.DoesInvalidTokenExist(ctx, token) {
			log.Printf("Log in system again.")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, appErrors.ErrInvalidToken.Error())
			return
		}

		if err := forger.Decode(token); err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenMalformed) || errors.Is(err, jwt.ErrTokenNotValidYet) {
				log.Printf("Invalid Token: %s", err)
				ctx.AbortWithStatusJSON(http.StatusBadRequest, appErrors.ErrInvalidToken.Error())
				return
			}
			log.Printf("Internal server error: %s", err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, "")
		}
		ctx.Next()
	}
}
