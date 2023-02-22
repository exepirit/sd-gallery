package leveldb

import (
	"context"
	"encoding/json"

	"github.com/exepirit/sd-gallery/internal/model"
	"github.com/exepirit/sd-gallery/internal/repository"
	"github.com/google/uuid"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/util"
)

// PictureRepository implements repository.PictureRepository uses LevelDB.
type PictureRepository struct {
	db *leveldb.DB
}

func (repo PictureRepository) GetOne(ctx context.Context, pictureId uuid.UUID) (model.Picture, error) {
	data, err := repo.db.Get([]byte("picture:"+pictureId.String()), nil)
	if err != nil {
		return model.Picture{}, err
	}

	var picture model.Picture
	err = json.Unmarshal(data, &picture)
	return picture, err
}

func (repo PictureRepository) Query(ctx context.Context) repository.Query[model.Picture] {
	iter := repo.db.NewIterator(util.BytesPrefix([]byte("picture:")), nil)
	return &PictureQuery{iter: iter, limit: -1}
}

func (repo PictureRepository) Put(ctx context.Context, picture model.Picture) error {
	data, err := json.Marshal(picture)
	if err != nil {
		return err
	}

	return repo.db.Put([]byte("picture:"+picture.ID.String()), data, nil)
}

func (repo PictureRepository) Delete(ctx context.Context, pictureId uuid.UUID) error {
	return repo.db.Delete([]byte("picture:"+pictureId.String()), nil)
}

// PictureQuery implements repository.Query[model.Picture] with LevelDB.
type PictureQuery struct {
	iter  iterator.Iterator
	skip  int
	limit int
}

func (q *PictureQuery) Skip(n int) {
	q.skip = n
}

func (q *PictureQuery) Limit(n int) {
	q.limit = n
}

func (q PictureQuery) GetAll() ([]model.Picture, error) {
	pictures := make([]model.Picture, 0)
	var picture model.Picture
	var skipped int

	for i := 0; q.iter.Next(); i++ {
		if q.limit >= 0 && len(pictures) == q.limit {
			break
		}

		if skipped < q.skip {
			skipped++
			continue
		}

		if err := json.Unmarshal(q.iter.Value(), &picture); err != nil {
			return pictures, err
		}
		pictures = append(pictures, picture)
	}
	q.iter.Release()
	return pictures, q.iter.Error()
}

func (q PictureQuery) Iterate(callee func(p model.Picture) bool) error {
	var picture model.Picture
	var skipped, returned int

	for q.iter.Next() {
		if q.limit >= 0 && returned == q.limit {
			break
		}

		if skipped < q.skip {
			skipped++
			continue
		}

		if err := json.Unmarshal(q.iter.Value(), &picture); err != nil {
			return err
		}

		if !callee(picture) {
			break
		}
	}
	q.iter.Release()
	return q.iter.Error()
}
