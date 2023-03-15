package leveldb

import (
	"encoding/json"
	"fmt"

	"github.com/exepirit/sd-gallery/internal/model"
	"github.com/exepirit/sd-gallery/internal/repository"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func newPictureQuery(db *leveldb.DB) *PictureQuery {
	return &PictureQuery{
		db:     db,
		limit:  -1,
		sortBy: "any",
	}
}

// PictureQuery implements repository.Query[model.Picture] with LevelDB.
type PictureQuery struct {
	db     *leveldb.DB
	limit  int
	skip   int
	sortBy string
}

func (q *PictureQuery) Skip(n int) repository.Query[model.Picture] {
	q.skip = n
	return q
}

func (q *PictureQuery) Limit(n int) repository.Query[model.Picture] {
	q.limit = n
	return q
}

func (q *PictureQuery) SortBy(key string) repository.Query[model.Picture] {
	q.sortBy = key
	return q
}

func (q *PictureQuery) prepareIterator() (iterable, error) {
	var iter iterable
	switch q.sortBy {
	case "any":
		stdIterator := q.db.NewIterator(util.BytesPrefix([]byte("picture:")), nil)
		iter = &basicIterator{
			Iterator: stdIterator,
		}
	case "ScrapeTime":
		index, err := getIndex(q.db, "Picture", q.sortBy)
		if err != nil {
			return nil, err
		}
		iter = &indexIterator{
			db:    q.db,
			index: index,
		}
	default:
		return nil, fmt.Errorf("unknown sort criteria %q", q.sortBy)
	}

	iter.Skip(q.skip)
	return iter, nil
}

func (q PictureQuery) GetAll() ([]model.Picture, error) {
	pictures := make([]model.Picture, 0)
	var picture model.Picture

	iter, err := q.prepareIterator()
	if err != nil {
		return pictures, err
	}
	defer iter.Release()

	for i := 0; iter.Next(); i++ {
		if q.limit >= 0 && len(pictures) == q.limit {
			break
		}

		if err := json.Unmarshal(iter.Value(), &picture); err != nil {
			return pictures, err
		}
		pictures = append(pictures, picture)
	}
	return pictures, iter.Error()
}

func (q PictureQuery) Iterate(callee func(p model.Picture) bool) error {
	var picture model.Picture
	var returned int

	iter, err := q.prepareIterator()
	if err != nil {
		return err
	}
	defer iter.Release()

	for iter.Next() {
		if q.limit >= 0 && returned == q.limit {
			break
		}

		if err := json.Unmarshal(iter.Value(), &picture); err != nil {
			return err
		}

		if !callee(picture) {
			break
		}
	}
	return iter.Error()
}
