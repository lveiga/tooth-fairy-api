package http

import (
	net "net/http"

	"github.com/gin-gonic/gin"
)

// BadRequest - returns when payload contains invalid data
func BadRequest(ctx *gin.Context, err error) {
	ctx.JSON(net.StatusBadRequest, gin.H{
		"code":    net.StatusBadRequest,
		"message": err.Error(),
	})
}

// NotFound - returns when client request a invalid route
func NotFound(ctx *gin.Context) {
	ctx.JSON(net.StatusNotFound, gin.H{
		"code":    net.StatusNotFound,
		"message": "rota não encontrada",
	})
}

// InternalServerError - returns when an error occurs in server
func InternalServerError(ctx *gin.Context) {
	ctx.JSON(net.StatusInternalServerError, gin.H{
		"code":    net.StatusInternalServerError,
		"message": "ocorreu um erro interno, favor tente mais tarde",
	})
}
