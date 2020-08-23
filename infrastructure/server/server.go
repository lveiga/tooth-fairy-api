package server

import (
	"context"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	ginlogrus "github.com/toorop/gin-logrus"
	"github.com/tooth-fairy/config"
	"github.com/tooth-fairy/docs"
	"github.com/tooth-fairy/infrastructure/database"
	h "github.com/tooth-fairy/infrastructure/http"
	"github.com/tooth-fairy/infrastructure/log"
	"github.com/tooth-fairy/infrastructure/middlewares"
	"go.elastic.co/apm/module/apmgin"
)

func getGinExecMode(c *config.AppConfig) string {
	if c.Environment == "local" {
		return gin.DebugMode
	}

	return gin.ReleaseMode
}

// Application - represents a application server configuration
type Application struct {
	config     *config.AppConfig
	database   *database.Database
	httpServer *http.Server
	router     *gin.Engine
	logger     log.Logger
}

// GetConfig - getter to retrieve server configuration
func (a *Application) GetConfig() *config.AppConfig {
	return a.config
}

// GetDatabase - getter to retreive database connection
func (a *Application) GetDatabase() *database.Database {
	return a.database
}

// WithMiddlewares - responsible to attach middlewares into http request pipeline
func (a *Application) WithMiddlewares() *Application {
	a.router.Use(cors.Default())
	a.router.Use(helmet.Default())
	a.router.Use(gzip.Gzip(gzip.BestCompression))
	a.router.Use(apmgin.Middleware(a.router, apmgin.WithRequestIgnorer(func(request *http.Request) bool {
		return request.URL.Path == "/health"
	})))

	a.router.Use(ginlogrus.Logger(a.logger.GetLogger()))

	a.router.NoRoute(func(ctx *gin.Context) {
		h.NotFound(ctx)
	})

	//TODO: CUSTOMIZE RECOVERY MIDDLEWARE AND CHECK APPLICATION PANICS
	a.router.Use(gin.Recovery())
	docs.SwaggerInfo.Title = "Tooth Fairy API"
	docs.SwaggerInfo.Description = "This is a server of Tooth Fairy."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "tooth-fairy.swagger.io"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	a.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return a
}

// WithHandlers - responsible to register application endpoints
func (a *Application) WithHandlers(routePrefix string, handlers ...Bindable) *Application {
	var router = a.router.Group(routePrefix)

	for _, handler := range handlers {
		handler.Bind(router, a)
	}

	return a
}

// Start ...
func (a *Application) Start() {
	var logger = a.logger.GetLogger()

	a.httpServer.Handler = a.router
	if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("listen: %s\n", err)
	}
}

// Shutdown ...
func (a *Application) Shutdown(ctx context.Context) error {
	return a.Shutdown(ctx)
}

// WithHealthcheck ...
func (a *Application) WithHealthcheck() *Application {
	a.router.GET("/health", middlewares.Health(a.database))
	return a
}

// New - responsible to creates a new instance from Application
func New(config *config.AppConfig, db *database.Database, router *gin.Engine, logger log.Logger) *Application {
	gin.SetMode(getGinExecMode(config))

	var server = &http.Server{
		Addr:    config.BindAddr,
		Handler: router,
	}

	var application = &Application{
		config:     config,
		database:   db,
		httpServer: server,
		router:     router,
		logger:     logger,
	}

	return application
}
