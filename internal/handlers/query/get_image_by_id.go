package query

import (
	"context"

	"github.com/exepirit/sd-gallery/internal/model"
	"github.com/exepirit/sd-gallery/internal/repository"
	"github.com/google/uuid"
)

type GetPictureByIDArgs struct {
	PictureID uuid.UUID
}

type GetPictureByIDResult model.Picture

// GetPictureByIDQuery find picture in repository by ID.
type GetPictureByIDQuery struct {
	pictures repository.Picture
}

func (query *GetPictureByIDQuery) Handle(ctx context.Context, args GetPictureByIDArgs) (GetPictureByIDResult, error) {
	picture, err := query.pictures.GetOne(ctx, args.PictureID)
	if err != nil {
		return GetPictureByIDResult{}, err
	}

	return GetPictureByIDResult(picture), nil
}
