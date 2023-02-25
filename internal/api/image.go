package api

import (
	"net/http"

	"github.com/exepirit/sd-gallery/internal/handlers/query"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ImageEndpoints provide API for access images.
type ImageEndpoints struct {
	getPictureById query.IQuery[query.GetPictureByIDArgs, query.GetPictureByIDResult]
}

func NewImageEndpoints(queries *query.Queries) *ImageEndpoints {
	return &ImageEndpoints{
		getPictureById: queries.GetPictureByID,
	}
}

func (e ImageEndpoints) Bind(r gin.IRoutes) {
	r.GET("/api/images/:id", e.GetFullByID)
}

func (e ImageEndpoints) GetFullByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	picture, err := e.getPictureById.Handle(ctx, query.GetPictureByIDArgs{PictureID: id})
	if err != nil {
		panic(err)
	}

	ctx.File(picture.ScrapeInfo.Path)
}
