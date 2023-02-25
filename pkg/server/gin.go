package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerConfig struct {
	Debug         bool
	ListenAddress string
}

func NewServer(cfg ServerConfig) Server {
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	ginHandler := gin.New()
	ginHandler.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	server := &http.Server{
		Addr:    cfg.ListenAddress,
		Handler: ginHandler,
	}

	return Server{
		Server: server,
	}
}

// Server is a wrapper around gin.Engine. Supports gracefull setup and shutdown.
type Server struct {
	*http.Server
}

func (srv *Server) Bind(bindable Bindable) {
	if bindable == nil {
		return
	}
	bindable.Bind(srv.Handler.(*gin.Engine))
}
