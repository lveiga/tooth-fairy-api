package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tooth-fairy/infrastructure/database"
)

//Health ....
func Health(db *database.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := db.CheckLiveness(); err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":     http.StatusInternalServerError,
				"database": "DOWN",
				"app":      "DOWN",
				"error":    err,
			})
			return
		}
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":     http.StatusOK,
			"database": "UP",
			"app":      "UP",
		})
		return
	}
}
