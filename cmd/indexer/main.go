package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/exepirit/sd-gallery/internal/config"
	"github.com/exepirit/sd-gallery/internal/index"
	"github.com/exepirit/sd-gallery/internal/infrastructure"
	"github.com/exepirit/sd-gallery/pkg/image"
	"go.uber.org/zap"
)

func main() {
	databaseURL := flag.String("database", "leveldb:./data", "Database connection string")
	inputDir := flag.String("input", "", "Input directory")
	debugFlag := flag.Bool("debug", false, "Enable debug output")
	flag.Parse()

	if inputDir == nil || *inputDir == "" {
		fmt.Println("Input directory required. Use --help for view parameters.")
		os.Exit(1)
	}

	var logger *zap.Logger
	var err error
	if *debugFlag {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction(zap.IncreaseLevel(zap.InfoLevel))
	}
	if err != nil {
		fmt.Println("Cannot setup logging:", err)
	}

	cfg := config.Config{
		DatabaseAddress: *databaseURL,
	}
	repositories, err := infrastructure.MakeRepositories(cfg)
	if err != nil {
		logger.Error("Cannot build repositories", zap.Error(err))
		os.Exit(127)
	}

	imageFinder := image.PNGFinder{
		Path: *inputDir,
	}
	indexer := index.NewIndexer(index.NewIndexerArgs{
		Pictures: repositories.Picture,
		Logger:   logger,
	})

	_, err = indexer.IndexFound(context.Background(), &imageFinder)
	if err != nil {
		logger.Error("Index error", zap.Error(err))
		os.Exit(1)
	}
}
