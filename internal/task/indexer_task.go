package task

import (
	"context"

	"github.com/exepirit/sd-gallery/internal/index"
	"github.com/exepirit/sd-gallery/pkg/bg"
	"github.com/exepirit/sd-gallery/pkg/image"
	"go.uber.org/zap"
)

func NewIndexGalleryTask(logger *zap.Logger, imageFinder image.Finder, indexer *index.Indexer) bg.Task {
	return func(ctx context.Context) error {
		n, err := indexer.IndexFound(ctx, imageFinder)
		if err != nil {
			logger.Error("Failed to index images", zap.Error(err))
		}

		if n > 0 {
			logger.Info("New images indexed", zap.Int("count", n))
		}

		return nil
	}
}
