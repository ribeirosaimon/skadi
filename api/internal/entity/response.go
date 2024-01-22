package entity

import "github.com/gin-gonic/gin"

func HandleSuccess(ctx *gin.Context, code int, body interface{}) {
	ctx.AbortWithStatusJSON(code, body)
}

func ThrowError(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(code, &erroHandler{Code: code, Message: message})
}

type erroHandler struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
