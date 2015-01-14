package reporter

import (
	"io"
)

type StepFunc func(...interface{}) (int, string)

type P interface {
	ID() string
}

type R interface {
	Bytes() []byte
	String() string
	Enter(P, StepFunc, ...interface{})
	Exit()
	Path() string
	Printf(string, ...interface{})
	Pipe() io.Writer

	Report(StepFunc, ...interface{})
	Success(StepFunc, ...interface{})
}