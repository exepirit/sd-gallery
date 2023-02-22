package leveldb

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/exepirit/sd-gallery/internal/model"
	"github.com/exepirit/sd-gallery/internal/repository"
	"github.com/google/uuid"
	"github.com/syndtr/goleveldb/leveldb"
)

func TestLevelDB(t *testing.T) {
	pictures := map[string]model.Picture{
		"pic1": {
			ID:   uuid.New(),
			Name: "pic1",
		},
		"pic2": {
			ID:   uuid.New(),
			Name: "pic2",
		},
	}

	db, err := leveldb.OpenFile("test_db.leveldb", nil)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer func() {
		os.RemoveAll("test_db.leveldb")
	}()

	var repo repository.Picture = &PictureRepository{db: db}

	for _, pic := range pictures {
		if err := repo.Put(context.TODO(), pic); err != nil {
			t.Fatalf("unexpected error while pushing value: %s", err)
		}
	}

	err = repo.Query(context.TODO()).
		Iterate(func(actual model.Picture) bool {
			expected, ok := pictures[actual.Name]
			if !ok {
				t.Fatalf("unexpected picture %#v", actual)
				return true
			}

			if !reflect.DeepEqual(actual, expected) {
				t.Fatalf("corrupted data\n\texpected: %#v\n\tgot: %#v", expected, actual)
			}
			return true
		})
	if err != nil {
		t.Fatalf("unexpected error while iterating: %s", err)
	}
}
