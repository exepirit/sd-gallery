package api

import (
	"net/http"
	"strconv"

	"github.com/exepirit/sd-gallery/internal/handlers"
	"github.com/exepirit/sd-gallery/internal/handlers/query"
	"github.com/gin-gonic/gin"
)

func NewFeedEndpoints(handlers *handlers.Handlers) *FeedEndpoints {
	return &FeedEndpoints{
		handlers: handlers,
	}
}

type FeedEndpoints struct {
	handlers *handlers.Handlers
}

func (e FeedEndpoints) Bind(routes gin.IRoutes) {
	routes.GET("/api/feed", e.GetFeedPage)
}

type ImageDto struct {
	PictureID string `json:"pictureId"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}

type ImagesPageDto struct {
	Items []ImageDto `json:"items"`
	Page  int        `json:"page"`
}

func (e FeedEndpoints) GetFeedPage(ctx *gin.Context) {
	pageIdx, err := strconv.Atoi(ctx.Query("page"))
	if err != nil || pageIdx < 0 {
		pageIdx = 0
	}

	pageSize, err := strconv.Atoi(ctx.Query("count"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	result, err := e.handlers.Queries.GetPictureFeed.Handle(ctx, query.GetPictureFeedArgs{
		PageSize:  pageSize,
		PageIndex: pageIdx,
	})
	if err != nil {
		panic(err)
	}

	response := ImagesPageDto{
		Items: make([]ImageDto, 0, len(result.Pictures)),
		Page:  pageIdx,
	}
	for _, picture := range result.Pictures {
		response.Items = append(response.Items, ImageDto{
			PictureID: picture.ID.String(),
			Name: picture.Name,
			URL:  "/api/images/" + picture.ID.String(),
			Width: picture.Size.Width,
			Height: picture.Size.Height,
		})
	}
	ctx.JSON(http.StatusOK, response)
}
