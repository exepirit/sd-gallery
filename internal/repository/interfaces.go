package repository

import (
	"context"

	"github.com/exepirit/sd-gallery/internal/model"
	"github.com/google/uuid"
)

type Picture interface {
	GetOne(ctx context.Context, pictureId uuid.UUID) (model.Picture, error)
	Query(ctx context.Context) Query[model.Picture]
	Put(ctx context.Context, picture model.Picture) error
	Delete(ctx context.Context, pictureId uuid.UUID) error
}

type Query[T any] interface {
	Skip(n int) Query[T]
	Limit(n int) Query[T]
	GetAll() ([]T, error)
	Iterate(callee func(m T) bool) error
}
