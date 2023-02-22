package model

import (
	"time"

	"github.com/google/uuid"
)

type Picture struct {
	ID           uuid.UUID
	Name         string
	Size         PictureSize
	ScrapeInfo   ScrapeInfo
	Tags         []string
	GenerateInfo GenerateInfo
}

type PictureSize struct {
	Width  int
	Height int
}

type ScrapeInfo struct {
	Time     time.Time
	Path     string
	FileHash string
}

type GenerateInfo struct {
	Prompt         string
	NegativePrompt string
	Steps          int
	Size           string
	Seed           string
	Sampler        string
	CfgScale       float64
	Model          GenerateModel
	Additional     map[string]string
}

type GenerateModel struct {
	Name string
	Hash string
}
