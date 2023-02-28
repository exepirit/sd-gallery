package index

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/exepirit/sd-gallery/pkg/image"
)

func RecognizeSdOutput(img image.Image) (PictureMeta, error) {
	var (
		metadata map[string]string
		err      error
	)
	switch img.Format {
	case "png":
		metadata, err = loadPngMeta(img.Path)
	default:
		return PictureMeta{}, errors.New("unknown format")
	}
	if err != nil {
		return PictureMeta{}, err
	}

	parameters, ok := metadata["parameters"]
	if !ok {
		return PictureMeta{}, ErrNotGeneratedBySD
	}

	output := PictureMeta{
		Image: img,
	}

	if ok := parseParameters(&output, parameters); !ok {
		return output, ErrNotGeneratedBySD
	}
	return output, nil
}

func parseParameters(output *PictureMeta, text string) bool {
	parts := strings.SplitN(text, "\n", 3)

	var other string
	switch len(parts) {
	case 2:
		output.Prompt = parts[0]
		other = parts[1]
	case 3:
		output.Prompt = parts[0]
		output.NegativePrompt = strings.TrimPrefix(parts[1], "Negative prompt: ")
		other = parts[2]
	}

	options := strings.Split(other, ", ")
	for _, opt := range options {
		optKV := strings.SplitN(opt, ": ", 2)
		key, value := optKV[0], optKV[1]
		switch key {
		case "Steps":
			output.Steps, _ = strconv.Atoi(value)
		case "Sampler":
			output.Sampler = value
		case "CFG scale":
			v, err := strconv.Atoi(value)
			if err != nil {
				output.CfgScale, _ = strconv.ParseFloat(key, 32)
			} else {
				output.CfgScale = float64(v)
			}
		case "Seed":
			output.Seed = value
		case "Size":
			output.Size = value
		case "Model hash":
			output.Model.Hash = value
		case "Model":
			output.Model.Name = value
		}
	}

	return true
}

func loadPngMeta(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	metadata := make(map[string]string, 1)
	chunks, err := image.DecodePngTextChunk(file)
	if err != nil {
		return metadata, err
	}

	for _, chunk := range chunks {
		if chunk.Name != "tEXt" {
			continue
		}

		keyword := readString(chunk.Data)
		text := string(chunk.Data[len(keyword)+2:])
		metadata[keyword] = text
	}

	return metadata, nil
}

func readString(data []byte) string {
	for i := 0; i < len(data); i++ {
		if data[i] == 0 {
			return string(data[:i])
		}
	}
	return ""
}
