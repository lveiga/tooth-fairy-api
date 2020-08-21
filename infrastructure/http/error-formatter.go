package http

import (
	net "net/http"

	"github.com/gin-gonic/gin"
)

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// BadRequest - returns when payload contains invalid data
func BadRequest(ctx *gin.Context, err error) {
	er := HTTPError{
		Code:    net.StatusBadRequest,
		Message: err.Error(),
	}
	ctx.JSON(net.StatusBadRequest, er)
}

// NotFound - returns when client request a invalid route
func NotFound(ctx *gin.Context) {
	er := HTTPError{
		Code:    net.StatusNotFound,
		Message: "rota não encontrada",
	}
	ctx.JSON(net.StatusNotFound, er)
}

// InternalServerError - returns when an error occurs in server
func InternalServerError(ctx *gin.Context) {
	er := HTTPError{
		Code:    net.StatusInternalServerError,
		Message: "rota não encontrada",
	}
	ctx.JSON(net.StatusInternalServerError, er)
}
