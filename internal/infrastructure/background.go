package infrastructure

import "github.com/exepirit/sd-gallery/pkg/bg"

func NewTaskRunner() *bg.TaskRunner {
	return &bg.TaskRunner{}
}

func NewBackgroundWorker(runner *bg.TaskRunner) *bg.BackgroundWorker {
	return &bg.BackgroundWorker{TaskRunner: runner}
}
