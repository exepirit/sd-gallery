package image

import (
	"io/fs"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

type Image struct {
	Name   string
	Path   string
	Format string
}

type Finder interface {
	Find() ([]Image, error)
}

type PNGFinder struct {
	Path      string
	Recursive bool
}

func (finder *PNGFinder) Find() ([]Image, error) {
	if finder.Recursive || true {
		images := make([]Image, 0)
		err := filepath.WalkDir(finder.Path, func(path string, d fs.DirEntry, err error) error {
			if !d.IsDir() {
				return nil
			}

			imgs, err := finder.findInDir(path)
			images = append(images, imgs...)
			return err
		})
		return images, err
	}

	return finder.findInDir(finder.Path)
}

func (finder *PNGFinder) findInDir(dirPath string) ([]Image, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	images := make([]Image, 0, len(entries))

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		entryPath := path.Join(dirPath, entry.Name())
		if finder.fileIsPNG(entryPath) {
			images = append(images, Image{
				Name:   entry.Name(),
				Path:   entryPath,
				Format: "png",
			})
		}
	}

	return images, nil

}

func (*PNGFinder) fileIsPNG(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()

	var buf [512]byte
	if _, err := file.Read(buf[:]); err != nil {
		return false
	}

	mimeType := http.DetectContentType(buf[:])
	return mimeType == "image/png"
}

func CombineFinders(finders ...Finder) Finder {
	return CombinedFinder(finders)
}

type CombinedFinder []Finder

func (finders CombinedFinder) Find() ([]Image, error) {
	images := make([]Image, 0)
	for _, finder := range finders {
		newImages, err := finder.Find()
		if err != nil {
			return images, err
		}

		images = append(images, newImages...)
	}
	return images, nil
}
