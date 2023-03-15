package query

import (
	"context"

	"github.com/exepirit/sd-gallery/internal/model"
	"github.com/exepirit/sd-gallery/internal/repository"
)

type GetPictureFeedArgs struct {
	PageSize  int
	PageIndex int
}

type GetPictureFeedResult struct {
	Pictures []model.Picture
}

// GetPictureFeedQuery fetch generated pictures feed.
type GetPictureFeedQuery struct {
	picturesRepo repository.Picture
}

func (q *GetPictureFeedQuery) Handle(ctx context.Context, args GetPictureFeedArgs) (GetPictureFeedResult, error) {
	pictures, err := q.picturesRepo.Query(ctx).
		SortBy("ScrapeTime").
		Skip(args.PageSize * args.PageIndex).
		Limit(args.PageSize).
		GetAll()

	return GetPictureFeedResult{
		Pictures: pictures,
	}, err
}
