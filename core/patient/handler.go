package patient

import (
	"github.com/gin-gonic/gin"
	"github.com/tooth-fairy/infrastructure/server"
)

//Handler - represents a route/controller binder
type Handler struct{}

//Bind - method responsible to bind controller and actions
func (h *Handler) Bind(router *gin.RouterGroup, app *server.Application) {
	var database = app.GetDatabase()
	var controller = newController(newRepository(database))

	router.POST("/patients", controller.NewPatient)
	router.GET("/patients", controller.GetPacients)
	router.GET("/patients:id", controller.GetPatient)
	router.PUT("/patients:id", controller.UpdatePatient)
	router.DELETE("/patients:id", controller.DeletePatient)
}
