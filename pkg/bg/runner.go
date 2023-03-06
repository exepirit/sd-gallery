package bg

import (
	"context"
	"errors"
)

// TaskRunner is a registry of task, accessed by their IDs.
type TaskRunner struct {
	tasks map[TaskID]Task
}

// Register puts task into registry.
func (runner *TaskRunner) Register(taskId TaskID, task Task) error {
	_, exists := runner.tasks[taskId]
	if exists {
		return errors.New("already registered")
	}

	runner.tasks[taskId] = task
	return nil
}

// Run starts the task immediately.
func (runner *TaskRunner) Run(taskId TaskID, ctx context.Context) error {
	task, ok := runner.tasks[taskId]
	if !ok {
		return errors.New("task not found")
	}

	return task(ctx)
}

// Run starts the task in new goroutine.
func (runner *TaskRunner) RunAsync(taskId TaskID, ctx context.Context) error {
	task, ok := runner.tasks[taskId]
	if !ok {
		return errors.New("task not found")
	}

	go func(ctx context.Context, task Task) {
		_ = task(ctx)
	}(ctx, task)

	return nil
}
