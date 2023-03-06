package bg

import (
	"context"
	"errors"
	"math"
	"time"
)

// BackgroundWorker manage plan of background tasks and run tasks by schedule.
type BackgroundWorker struct {
	schedule    map[TaskID]time.Duration
	lastPlanned map[TaskID]time.Time

	interruptChan chan struct{}
	runner        *TaskRunner
}

// Schedule put task into execute schedule.
func (worker *BackgroundWorker) Schedule(taskID TaskID, task Task, period time.Duration) error {
	err := worker.runner.Register(taskID, task)
	if err != nil {
		return err
	}

	// TODO: thread-safe zone
	worker.schedule[taskID] = period
	worker.lastPlanned[taskID] = time.Unix(0, 0)

	return nil
}

// Run starts background worker infinite routine. Blocking function.
func (worker *BackgroundWorker) Run() {
	// init interrupt channel
	worker.interruptChan = make(chan struct{})

	for {
		// some calculations
		nextTaskId, waitTime := worker.getNextTask()
		timer := time.NewTimer(waitTime)

		// create new cancellable context for task
		ctx, cancelTask := context.WithCancel(context.Background())

		// wait for scheduled task or interrupt
		select {
		case <-timer.C:
			err := worker.runner.Run(nextTaskId, ctx)
			if err != nil {
				panic(err) // TODO: log errors
			}
		case <-worker.interruptChan:
			cancelTask()
			return
		}
	}
}

// Stop interrupt routine of this worker.
func (worker *BackgroundWorker) Stop() error {
	if worker.interruptChan == nil {
		return errors.New("scheduler not running")
	}

	close(worker.interruptChan)
	return nil
}

func (worker *BackgroundWorker) getNextTask() (TaskID, time.Duration) {
	now := time.Now()

	var nextTask TaskID
	var waitTime time.Duration = time.Duration(math.MaxInt64)

	for taskId, period := range worker.schedule { // TODO: thread-safe read
		lastRunAgo := now.Sub(worker.lastPlanned[taskId])
		wait := period - lastRunAgo

		if wait < waitTime {
			nextTask = taskId
			waitTime = wait
		}
	}

	if waitTime < 0 {
		waitTime = 0
	}

	worker.lastPlanned[nextTask] = now
	return nextTask, waitTime
}
