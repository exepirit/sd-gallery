package bg

import "context"

// Task describe type of function of any task.
type Task func(ctx context.Context) error

// TaskID is a wrapper around string for task identification purposes.
type TaskID string
