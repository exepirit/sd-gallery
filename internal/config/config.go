package config

// Config contains whole application configuration.
type Config struct {
	// ListenAddress is an address, which HTTP listen to.
	ListenAddress string `mapstructure:"LISTEN_ADDRESS"`

	// DatabaseAddress is a connection to database string.
	// Format is URL-like, which schema defines DB type.
	// Example: leveldb://folder
	DatabaseAddress string `mapstructure:"DATABASE_ADDRESS"`

	// Index contains configuration for image indexer.
	Index IndexerConfig `mapstructure:"INDEX"`
}

// IndexerConfig contains image indexer config.
type IndexerConfig struct {
	// Paths is a enumeration of paths, which images could be find.
	Paths []string `mapstructure:"PATHS"`
}

// Default make default application configuration.
func Default() Config {
	return Config{
		ListenAddress:   "[::]:8080",
		DatabaseAddress: "leveldb:./data",
		Index: IndexerConfig{
			Paths: []string{},
		},
	}
}
