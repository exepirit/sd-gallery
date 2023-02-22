package config

// Config contains whole application configuration.
type Config struct {
	// ListenAddress is an address, which HTTP listen to.
	ListenAddress string

	// DatabaseAddress is a connection to database string.
	// Format is URL-like, which schema defines DB type.
	// Example: leveldb://folder
	DatabaseAddress string

	// Index contains configuration for image indexer.
	Index IndexerConfig
}

// IndexerConfig contains image indexer config.
type IndexerConfig struct {
	// Paths is a enumeration of paths, which images could be find.
	Paths []string
}

// Default make default application configuration.
func Default() Config {
	return Config{
		ListenAddress:   "[::]:8080",
		DatabaseAddress: "leveldb://./data",
		Index: IndexerConfig{
			Paths: []string{},
		},
	}
}
