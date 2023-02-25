// This module provide REST API for service.

package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewFeedEndpoints),
	fx.Provide(NewImageEndpoints),
	fx.Provide(NewAPI),
)

func NewAPI(feed *FeedEndpoints, image *ImageEndpoints) *API {
	return &API{
		Feed: feed,
		Image: image,
	}
}

type API struct {
	Feed *FeedEndpoints
	Image *ImageEndpoints
}

func (api API) Bind(r gin.IRoutes) {
	api.Feed.Bind(r)
	api.Image.Bind(r)
}
