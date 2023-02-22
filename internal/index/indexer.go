package index

import (
	"context"
	"errors"
	"fmt"

	"github.com/exepirit/sd-gallery/internal/model"
	"github.com/exepirit/sd-gallery/internal/repository"
	"github.com/exepirit/sd-gallery/pkg/image"
)

// NewIndexer makes new Indexer instance.
func NewIndexer(picturesRepo repository.Picture) Indexer {
	return Indexer{
		IgnoreErrors: true,
		picturesRepo: picturesRepo,
	}
}

// Indexer store new pictures data in repository.
type Indexer struct {
	IgnoreErrors bool

	picturesRepo repository.Picture
	indexCache   map[string]struct{}
}

// AddToIndex scrape additional data and add picture in repository.
func (indexer *Indexer) AddToIndex(ctx context.Context, image image.Image) error {
	if indexer.indexCache == nil {
		indexer.fillIndexCache(ctx)
	}

	fileHash, err := calcFileHash(image.Path)
	if err != nil {
		return fmt.Errorf("cannot calculate file hash: %w", err)
	}

	if _, inCache := indexer.indexCache[fileHash]; inCache {
		return errors.New("already indexed")
	}

	pictureMeta, err := RecognizeSdOutput(image)
	if err != nil {
		return fmt.Errorf("cannot interpret image as Stable Diffusion output: %w", err)
	}

	picture := newPicture(pictureMeta, fileHash)

	if err = indexer.picturesRepo.Put(ctx, picture); err != nil {
		return fmt.Errorf("cannot store picture data: %w", err)
	}
	return nil
}

// IndexFound index in repository every found picture.
func (indexer Indexer) IndexFound(ctx context.Context, finder image.Finder) (int, error) {
	found, err := finder.Find()
	if err != nil {
		return 0, err
	}

	counter := 0
	for _, image := range found {
		if err := indexer.AddToIndex(ctx, image); err != nil {
			if indexer.IgnoreErrors {
				continue
			}
			return counter, err
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
