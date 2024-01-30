package entity

import "github.com/gin-gonic/gin"

func HandleSuccess(ctx *gin.Context, code int, body interface{}) {
	ctx.AbortWithStatusJSON(code, body)
}

func ThrowError(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(code, &erroHandler{Code: code, Message: message})
}

func AuthorizationError(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(401, &erroHandler{Code: 401, Message: "This user was not authenticated"})
}

type erroHandler struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
