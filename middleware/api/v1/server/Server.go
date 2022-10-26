package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-errors/errors"

	"github.com/gin-gonic/gin"
	v1 "github.com/kaan-devoteam/one-click-deploy-demo/api/v1"
	"github.com/kaan-devoteam/one-click-deploy-demo/log"
	"github.com/kaan-devoteam/one-click-deploy-demo/settings"
)

type Server struct {
	activeServer *http.Server
}

func New() *Server {
	return &Server{}
}

func (s *Server) Run() {
	name := fmt.Sprintf("%s %s", settings.Name, settings.Version)
	s.Init()
	log.Info("%s listening to port %d", name, settings.ServerPort)

	err := s.activeServer.ListenAndServe()
	if err != nil {
		log.Error("%s stopped with error\n%s", name, errors.New(err).ErrorStack())
	} else {
		log.Info("%s cache server closed", name)
	}
}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := s.activeServer.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Init() {
	s.activeServer = &http.Server{
		Addr:              fmt.Sprintf(":%d", settings.ServerPort),
		Handler:           App(),
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      settings.ServerTimeout,
		IdleTimeout:       settings.ServerTimeout,
	}
}

func App() *gin.Engine {
	// setGinMode() // TODO: Make sure gin mode comes from the environment
	log.Info("Gin mode is set to %v", gin.Mode())
	app := gin.New()
	app.Use(
		SetServerHeader(fmt.Sprintf("%s/%s", settings.Name, settings.Version)),
		gin.Recovery(),
	)
	if settings.ServerAccessLog {
		app.Use(gin.Logger())
	}
	v1.Router(app)
	return app
}

func SetServerHeader(name string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Header("Server", name)
		ctx.Next()
	}
}
