package index

import (
	"context"
	"fmt"

	"github.com/exepirit/sd-gallery/internal/model"
	"github.com/exepirit/sd-gallery/internal/repository"
	"github.com/exepirit/sd-gallery/pkg/image"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type NewIndexerArgs struct {
	fx.In

	Pictures repository.Picture
	Logger   *zap.Logger `optional:"true"`
}

// NewIndexer makes new Indexer instance.
func NewIndexer(dependencies NewIndexerArgs) Indexer {
	if dependencies.Logger == nil {
		dependencies.Logger, _ = zap.NewDevelopment()
	}

	return Indexer{
		logger: dependencies.Logger,
		picturesRepo: dependencies.Pictures,
	}
}

// Indexer store new pictures data in repository.
type Indexer struct {
	logger *zap.Logger
	picturesRepo repository.Picture
	indexCache   map[string]struct{}
}

// addToIndex scrape additional data and add picture in repository.
func (indexer *Indexer) addToIndex(ctx context.Context, image image.Image) error {
	if indexer.indexCache == nil {
		indexer.fillIndexCache(ctx)
	}

	fileHash, err := calcFileHash(image.Path)
	if err != nil {
		return fmt.Errorf("cannot calculate file hash: %w", err)
	}

	if _, inCache := indexer.indexCache[fileHash]; inCache {
		return ErrAlreadyIndexed
	}

	pictureMeta, err := RecognizeSdOutput(image)
	if err != nil {
		return ErrNotGeneratedBySD
	}

	picture := newPicture(pictureMeta, fileHash)

	if err = indexer.picturesRepo.Put(ctx, picture); err != nil {
		return fmt.Errorf("cannot store picture data: %w", err)
	}

	indexer.logger.Info("Picture indexed",
		zap.String("filename", image.Name),
		zap.String("uuid", picture.ID.String()))
	return nil
}

// IndexFound index in repository every found picture.
func (indexer Indexer) IndexFound(ctx context.Context, finder image.Finder) (int, error) {
	found, err := finder.Find()
	if err != nil {
		return 0, err
	}
	indexer.logger.Debug("Images scan completed", zap.Int("found", len(found)))

	counter := 0
	for _, image := range found {
		if err := indexer.addToIndex(ctx, image); err != nil {
			switch err {
			case ErrAlreadyIndexed:
				indexer.logger.Debug("File already in index", zap.String("filename", image.Name))
			case ErrNotGeneratedBySD:
				indexer.logger.Warn("Image not recognized as Stable Diffusion output",
					zap.String("filepath", image.Path))
			default:
				return counter, err
			}
		}
		counter++
	}

	return counter, nil
}

func (indexer *Indexer) fillIndexCache(ctx context.Context) error {
	indexer.indexCache = make(map[string]struct{})
	return indexer.picturesRepo.
		Query(ctx).
		Iterate(func(m model.Picture) bool {
			indexer.indexCache[m.ScrapeInfo.FileHash] = struct{}{}
			return true
		})
}
