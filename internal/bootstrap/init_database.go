package bootstrap

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/exepirit/sd-gallery/internal/config"
	"github.com/exepirit/sd-gallery/internal/repository"
	leveldbrepo "github.com/exepirit/sd-gallery/internal/repository/leveldb"
	"github.com/syndtr/goleveldb/leveldb"
)

// MakeRepositories create new repository connectors.
func MakeRepositories(cfg config.Config) (repository.Repositories, error) {
	connection, err := url.Parse(cfg.DatabaseAddress)
	if err != nil {
		return repository.Repositories{}, errors.New("invalid database connection string")
	}

	switch connection.Scheme {
	case "leveldb":
		return makeLevelDbRepositories(connection.Opaque)
	default:
		return repository.Repositories{}, fmt.Errorf("unknown database type %s", connection.Scheme)
	}
}

func makeLevelDbRepositories(dirname string) (repository.Repositories, error) {
	db, err := leveldb.OpenFile(dirname, nil)
	if err != nil {
		return repository.Repositories{}, fmt.Errorf("cannot open DB: %w", err)
	}

	return repository.Repositories{
		Picture: leveldbrepo.NewPictureRepository(db),
	}, nil
}
