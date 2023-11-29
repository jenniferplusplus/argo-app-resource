package concourse

import "io"

type Task struct {
	stdin  io.Reader
	stderr io.Writer
	stdout io.Writer
	args   []string
}

func NewTask(
	stdin io.Reader,
	stderr io.Writer,
	stdout io.Writer,
	args []string,
) *Task {
	return &Task{
		stdin:  stdin,
		stderr: stderr,
		stdout: stdout,
		args:   args,
	}
}
