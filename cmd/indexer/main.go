package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/exepirit/sd-gallery/internal/bootstrap"
	"github.com/exepirit/sd-gallery/internal/config"
	"github.com/exepirit/sd-gallery/internal/index"
	"github.com/exepirit/sd-gallery/pkg/image"
)

func main() {
	databaseURL := flag.String("database", "leveldb:./data", "Database connection string")
	inputDir := flag.String("input", "", "Input directory")
	flag.Parse()

	if inputDir == nil || *inputDir == "" {
		fmt.Println("Input directory required. Use --help for view parameters.")
		os.Exit(1)
	}

	cfg := config.Config{
		DatabaseAddress: *databaseURL,
	}
	repositories, err := bootstrap.MakeRepositories(cfg)
	if err != nil {
		fmt.Println("Cannot init storage:", err)
		os.Exit(127)
	}

	imageFinder := image.PNGFinder{
		Path: *inputDir,
	}
	indexer := index.NewIndexer(repositories.Picture)

	indexedCount, err := indexer.IndexFound(context.Background(), &imageFinder)
	if err != nil {
		fmt.Println("Indexing error:", err)
		os.Exit(1)
	}
	fmt.Printf("%d images indexed\n", indexedCount)
}
