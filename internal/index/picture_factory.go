package index

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/exepirit/sd-gallery/internal/model"
	"github.com/google/uuid"
)

func newPicture(pictureMeta PictureMeta, fileHash string) model.Picture {
	splitPrompt := strings.Split(pictureMeta.Prompt, ",")
	for i := 0; i < len(splitPrompt); i++ {
		splitPrompt[i] = strings.Trim(splitPrompt[i], " ")
	}

	var width, height int
	splitSize := strings.SplitN(pictureMeta.Size, "x", 2)
	if len(splitSize) == 2 {
		width, _ = strconv.Atoi(splitSize[0])
		height, _ = strconv.Atoi(splitSize[1])
	}

	return model.Picture{
		ID:   uuid.New(),
		Name: pictureMeta.Image.Name,
		Size: model.PictureSize{
			Width:  width,
			Height: height,
		},
		ScrapeInfo: model.ScrapeInfo{
			Time:     time.Now(),
			Path:     pictureMeta.Image.Path,
			FileHash: fileHash,
		},
		Tags: splitPrompt,
		GenerateInfo: model.GenerateInfo{
			Prompt:         pictureMeta.Prompt,
			NegativePrompt: pictureMeta.NegativePrompt,
			Steps:          pictureMeta.Steps,
			Size:           pictureMeta.Size,
			Seed:           pictureMeta.Seed,
			Sampler:        pictureMeta.Sampler,
			CfgScale:       pictureMeta.CfgScale,
			Model: model.GenerateModel{
				Name: pictureMeta.Model.Name,
				Hash: pictureMeta.Model.Hash,
			},
			Additional: map[string]string{},
		},
	}
}

func calcFileHash(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, f)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
