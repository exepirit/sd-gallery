package infrastructure

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewTaskRunner),
	fx.Provide(NewBackgroundWorker),
	fx.Provide(NewLogger),
	fx.Provide(MakeRepositories),
)
