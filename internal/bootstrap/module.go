package bootstrap

import (
	"github.com/exepirit/sd-gallery/internal/api"
	"github.com/exepirit/sd-gallery/internal/config"
	"github.com/exepirit/sd-gallery/internal/handlers"
	"github.com/exepirit/sd-gallery/internal/infrastructure"
	"github.com/exepirit/sd-gallery/internal/task"
	"github.com/exepirit/sd-gallery/pkg/server"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(func() config.Config {
		return config.MustLoad()
	}),
	fx.Provide(func(cfg config.Config) server.Server {
		return server.NewServer(server.ServerConfig{
			Debug:         true,
			ListenAddress: cfg.ListenAddress,
		})
	}),
	infrastructure.Module,
	handlers.Module,
	api.Module,
	task.Module,
	fx.Invoke(InitAppLifecycle),
)
