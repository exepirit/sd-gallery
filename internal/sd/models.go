package sd

import "github.com/exepirit/sd-gallery/internal/image"

type Output struct {
	Image          image.Image
	Prompt         string
	NegativePrompt string
	Steps          int
	Sampler        string
	CfgScale       float64
	Seed           string
	Size           string
	Model          Model
}

type Model struct {
	Name string
	Hash string
}
