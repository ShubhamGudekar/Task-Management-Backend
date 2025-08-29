package errors

import "errors"

var ErrTaskNotFound error = errors.New("task not found")
var ErrInvalidTaskStatus error = errors.New("invalid task status")
var ErrInvalidTaskPriority error = errors.New("invalid task priority")
