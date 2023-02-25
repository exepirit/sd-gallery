package query

import (
	"context"

	"github.com/exepirit/sd-gallery/internal/repository"
	"github.com/google/uuid"
)

type IQuery[TA any, TR any] interface {
	Handle(context.Context, TA) (TR, error)
}

type ICommand[TA any] interface {
	Handle(context.Context, TA) (uuid.UUID, error)
}

// Queries is a summary container for registered queries.
type Queries struct {
	GetPictureFeed IQuery[GetPictureFeedArgs, GetPictureFeedResult]
	GetPictureByID IQuery[GetPictureByIDArgs, GetPictureByIDResult]
}

// NewQueries initialize queries container.
func NewQueries(repo repository.Repositories) *Queries {
	return &Queries{
		GetPictureFeed: &GetPictureFeedQuery{picturesRepo: repo.Picture},
		GetPictureByID: &GetPictureByIDQuery{pictures: repo.Picture},
	}
}
