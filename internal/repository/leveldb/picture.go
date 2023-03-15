package leveldb

import (
	"context"
	"encoding/json"

	"github.com/exepirit/sd-gallery/internal/model"
	"github.com/exepirit/sd-gallery/internal/repository"
	"github.com/google/uuid"
	"github.com/syndtr/goleveldb/leveldb"
)

const pictureEntityName = "picture"

// NewPictureRepository create new PictureRepository instance.
func NewPictureRepository(db *leveldb.DB) *PictureRepository {
	return &PictureRepository{db: db}
}

// PictureRepository implements repository.PictureRepository uses LevelDB.
type PictureRepository struct {
	db *leveldb.DB
}

func (repo PictureRepository) GetOne(ctx context.Context, pictureId uuid.UUID) (model.Picture, error) {
	data, err := repo.db.Get(fmtEntityKey(pictureEntityName, pictureId), nil)
	if err != nil {
		return model.Picture{}, err
	}

	var picture model.Picture
	err = json.Unmarshal(data, &picture)
	return picture, err
}

func (repo PictureRepository) Query(ctx context.Context) repository.Query[model.Picture] {
	return newPictureQuery(repo.db)
}

func (repo PictureRepository) Put(ctx context.Context, picture model.Picture) error {
	data, err := json.Marshal(picture)
	if err != nil {
		return err
	}

	err = repo.db.Put(fmtEntityKey(pictureEntityName, picture.ID), data, nil)
	if err != nil {
		return err
	}

	index, err := getIndex(repo.db, "Picture", "ScrapeTime")
	if err != nil {
		return err
	}
	index.insert(
		fmtEntityKey(pictureEntityName, picture.ID),
		int(picture.ScrapeInfo.Time.UnixMilli()),
	)
	return putIndex(repo.db, "Picture", "ScrapeTime", index)
}

func (repo PictureRepository) Delete(ctx context.Context, pictureId uuid.UUID) error {
	index, err := getIndex(repo.db, "Picture", "ScrapeTime")
	if err != nil {
		return err
	}
	index.delete(fmtEntityKey(pictureEntityName, pictureId))
	if err = putIndex(repo.db, "Picture", "ScrapeTime", index); err != nil {
		return err
	}

	return repo.db.Delete(fmtEntityKey(pictureEntityName, pictureId), nil)
}
