package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tooth-fairy/config"
	"github.com/tooth-fairy/core/migration"
	"github.com/tooth-fairy/core/patient"
	"github.com/tooth-fairy/infrastructure/database"
	"github.com/tooth-fairy/infrastructure/log"
	"github.com/tooth-fairy/infrastructure/server"
)

func main() {
	var engine = gin.New()
	var config = config.New()
	var loggerAdapter = log.NewLogger(config.Environment)
	var logger = loggerAdapter.GetLogger()

	database, err := database.New(config)

	if err != nil {
		panic(err)
	}

	migration := migration.New(database)
	migration.Migrations()

	defer database.Close()

	var server = server.New(config, database, engine, loggerAdapter).
		WithMiddlewares().
		WithHealthcheck().
		WithHandlers("/api/v1",
			&patient.Handler{},
		)

	logger.Println("Start tooth fairy...")
	server.Start()
}
