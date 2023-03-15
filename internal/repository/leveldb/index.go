package leveldb

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
)

type index struct {
	Records []indexRecord `json:"records"`
}

type indexRecord struct {
	K []byte
	V int
}

func (r indexRecord) MarshalJSON() ([]byte, error) {
	s := struct {
		K string `json:"k"`
		V int
	}{
		K: base64.StdEncoding.EncodeToString(r.K),
		V: r.V,
	}
	return json.Marshal(&s)
}

func (r *indexRecord) UnmarshalJSON(data []byte) error {
	s := struct {
		K string `json:"k"`
		V int
	}{}
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	r.K, err = base64.StdEncoding.DecodeString(s.K)
	r.V = s.V
	return err
}

func (idx *index) insert(key []byte, value int) {
	newRecord := indexRecord{K: key, V: value}
	insertPos := 0
	for i, record := range idx.Records {
		if value < record.V {
			break
		}
		insertPos = i
	}

	if len(idx.Records) == 0 || len(idx.Records) == insertPos+1 {
		idx.Records = append(idx.Records, newRecord)
	} else {
		idx.Records = append(append(idx.Records[:insertPos], newRecord), idx.Records[insertPos:]...)
	}
}

func (idx *index) delete(key []byte) {
	for i, record := range idx.Records {
		if compareSlices(record.K, key) {
			switch {
			case len(idx.Records) == i+1:
				idx.Records = idx.Records[:i]
			case i == 0:
				idx.Records = idx.Records[i:]
			default:
				idx.Records = append(idx.Records[:i], idx.Records[i+1:]...)
			}
			return
		}
	}
}

type iterable interface {
	Next() bool
	Error() error
	Key() []byte
	Value() []byte
	Skip(n int)
	Release()
}

type indexIterator struct {
	db    *leveldb.DB
	index *index
	ptr   int
	err   error
	key   []byte
	value []byte
}

func (iter *indexIterator) Next() bool {
	if len(iter.index.Records) <= iter.ptr {
		return false
	}

	key := iter.index.Records[iter.ptr].K
	value, err := iter.db.Get(key, nil)
	if err != nil {
		iter.err = err
		return false
	}

	iter.ptr++
	iter.key, iter.value = key, value
	return true
}

func (iter indexIterator) Error() error {
	return iter.err
}

func (iter indexIterator) Key() []byte {
	return iter.key
}

func (iter indexIterator) Value() []byte {
	return iter.value
}

func (iter *indexIterator) Skip(n int) {
	if n >= len(iter.index.Records) || n < 0 {
		return
	}
	iter.ptr = n
}

func (iter *indexIterator) Release() {}

func compareSlices(s1, s2 []byte) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

type basicIterator struct {
	iterator.Iterator
	skip    int
	skipped int
}

func (iter *basicIterator) Next() bool {
	for iter.skipped < iter.skip {
		if !iter.Iterator.Next() {
			return false
		}
		iter.skipped++
	}

	return iter.Iterator.Next()
}

func (iter *basicIterator) Skip(n int) {
	if n >= 0 {
		iter.skip = n
	}
}

func (iter *basicIterator) Release() {
	iter.Iterator.Release()
}

func fmtEntityKey(entityName string, id uuid.UUID) []byte {
	return []byte(entityName + "_" + id.String())
}

func getIndex(db *leveldb.DB, entity, indexName string) (*index, error) {
	data, err := db.Get([]byte("INDEX."+entity+"."+indexName), nil)
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			return &index{
				Records: []indexRecord{},
			}, nil
		}
		return nil, err
	}

	var idx index
	err = json.Unmarshal(data, &idx)
	return &idx, err
}

func putIndex(db *leveldb.DB, entity, indexName string, idx *index) error {
	data, err := json.Marshal(idx)
	if err != nil {
		return err
	}

	return db.Put([]byte("INDEX."+entity+"."+indexName), data, nil)
}
