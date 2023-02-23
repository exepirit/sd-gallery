package handlers

import (
	"github.com/exepirit/sd-gallery/internal/handlers/query"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(query.NewQueries),
	fx.Provide(func(queries *query.Queries) *Handlers {
		return &Handlers{Queries: queries}
	}),
)

type Handlers struct {
	Queries *query.Queries
}
