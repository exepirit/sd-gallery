package main

import (
	"github.com/exepirit/sd-gallery/internal/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(bootstrap.Module).Run()
}
