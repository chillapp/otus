package app

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"otus/cmd/app/dependencies"
	health "otus/modules/common/health/delivery/http"
	users "otus/modules/common/users/delivery/http"
	"otus/storage"
)

type App struct {
	httpServer *http.Server
	Modules    *dependencies.Modules
}

func Initialize(store storage.Storage) *App {
	app := &App{}
	app.Modules = dependencies.Injection(store)
	app.httpServer = app.initHttpServer()
	return app
}

func (a *App) initHttpServer() *http.Server {
	router := a.regRouters()
	httpServer := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    5 * time.Minute,
		WriteTimeout:   5 * time.Minute,
		IdleTimeout:    90 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return httpServer
}

func (a *App) regRouters() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	health.RegHealthHTTP(router, a.Modules.Health)
	users.RegUsersHTTP(router, a.Modules.Users)
	return router
}

func (a *App) Run() error {
	return a.httpServer.ListenAndServe()
}

func (a *App) Shutdown(context context.Context) error {
	return a.httpServer.Shutdown(context)
}
