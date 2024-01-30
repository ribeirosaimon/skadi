package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/skadi/api/internal/entity"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			entity.AuthorizationError(ctx)
		}
	}
}
