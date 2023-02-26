// This module provide REST API for service.

package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewFeedEndpoints),
	fx.Provide(NewImageEndpoints),
	fx.Provide(NewPictureEndpoints),
	fx.Provide(NewAPI),
)

func NewAPI(
	feed *FeedEndpoints,
	image *ImageEndpoints,
	picture *PictureEndpoints,
) *API {
	return &API{
		Feed: feed,
		Image: image,
		Picture: picture,
	}
}

type API struct {
	Feed *FeedEndpoints
	Image *ImageEndpoints
	Picture *PictureEndpoints
}

func (api API) Bind(r gin.IRoutes) {
	api.Feed.Bind(r)
	api.Image.Bind(r)
	api.Picture.Bind(r)
}
