package server

import "github.com/gin-gonic/gin"

// Bindable - is an abstraction to bind resource handlers to definition
type Bindable interface {
	Bind(router *gin.RouterGroup, app *Application)
}
