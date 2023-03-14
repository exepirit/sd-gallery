package task

import (
	"time"

	"github.com/exepirit/sd-gallery/pkg/bg"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(fx.Annotate(NewIndexGalleryTask, fx.ResultTags(`name:"IndexGallery"`))),

	fx.Decorate(func(worker *bg.BackgroundWorker, tasks tasks) *bg.BackgroundWorker {
		tasks.register(worker)
		return worker
	}),
)

type tasks struct {
	fx.In

	IndexerTask bg.Task `name:"IndexGallery"`
}

func (tasks tasks) register(worker *bg.BackgroundWorker) {
	_ = worker.Schedule("IndexGallery", tasks.IndexerTask, time.Minute)
}
