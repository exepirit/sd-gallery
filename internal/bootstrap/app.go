package bootstrap

import (
	"context"
	"log"
	"net/http"

	"github.com/exepirit/sd-gallery/internal/api"
	"github.com/exepirit/sd-gallery/pkg/server"
	"go.uber.org/fx"
)

// InitAppLifecycle setup whole application lifecycle.
func InitAppLifecycle(
	srv server.Server,
	api *api.API,
	lifecycle fx.Lifecycle,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			srv.Bind(api)
			go func(server server.Server) {
				log.Printf("Handle new connections on %s", server.Addr)
				switch err := server.ListenAndServe(); err {
				case http.ErrServerClosed:
					return
				case nil:
					return
				default:
					panic(err)
				}
			}(srv)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			_ = srv.Shutdown(ctx)

			return nil
		},
	})
}
