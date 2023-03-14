package api

import (
	"net/http"

	"github.com/exepirit/sd-gallery/internal/handlers"
	"github.com/exepirit/sd-gallery/internal/handlers/query"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PictureEndpoints struct {
	getPictureByID query.IQuery[query.GetPictureByIDArgs, query.GetPictureByIDResult]
}

func NewPictureEndpoints(handlers *handlers.Handlers) *PictureEndpoints {
	return &PictureEndpoints{
		getPictureByID: handlers.Queries.GetPictureByID,
	}
}

func (e PictureEndpoints) Bind(r gin.IRoutes) {
	r.GET("/api/pictures/:id", e.GetByID)
}

type PictureDto struct {
	ID   uuid.UUID              `json:"id"`
	Name string                 `json:"name"`
	Size PictureSizeDto         `json:"size"`
	Tags []string               `json:"tags"`
	Info PictureGenerateInfoDto `json:"info"`
}

type PictureSizeDto struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type PictureGenerateInfoDto struct {
	Prompt         string  `json:"prompt"`
	NegativePrompt string  `json:"negativePrompt"`
	Steps          int     `json:"steps"`
	Size           string  `json:"size"`
	Seed           string  `json:"seed"`
	Sampler        string  `json:"sampler"`
	CfgScale       float32 `json:"cfgScale"`
	ModelName      string  `json:"modelName"`
	ModelHash      string  `json:"modelHash"`
}

func (e PictureEndpoints) GetByID(ctx *gin.Context) {
	pictureId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	picture, err := e.getPictureByID.Handle(ctx, query.GetPictureByIDArgs{
		PictureID: pictureId,
	})
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, PictureDto{
		ID:   picture.ID,
		Name: picture.Name,
		Size: PictureSizeDto{
			Width:  picture.Size.Width,
			Height: picture.Size.Height,
		},
		Tags: picture.Tags,
		Info: PictureGenerateInfoDto{
			Prompt:         picture.GenerateInfo.Prompt,
			NegativePrompt: picture.GenerateInfo.NegativePrompt,
			Steps:          picture.GenerateInfo.Steps,
			Size:           picture.GenerateInfo.Size,
			Seed:           picture.GenerateInfo.Seed,
			Sampler:        picture.GenerateInfo.Sampler,
			CfgScale:       float32(picture.GenerateInfo.CfgScale),
			ModelName:      picture.GenerateInfo.Model.Name,
			ModelHash:      picture.GenerateInfo.Model.Hash,
		},
	})
}
