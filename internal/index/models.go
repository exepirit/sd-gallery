package index

import "github.com/exepirit/sd-gallery/pkg/image"

type PictureMeta struct {
	Image          image.Image
	Prompt         string
	NegativePrompt string
	Steps          int
	Sampler        string
	CfgScale       float64
	Seed           string
	Size           string
	Model          StableDiffusionModel
}

type StableDiffusionModel struct {
	Name string
	Hash string
}
