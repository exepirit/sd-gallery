package index

import "errors"

var (
	ErrAlreadyIndexed = errors.New("already indexed")
	ErrNotGeneratedBySD = errors.New("cannot interpret file as Stable Diffusion output")
)
