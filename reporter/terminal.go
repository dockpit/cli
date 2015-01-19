package reporter

import (
	"bufio"
	"bytes"
	"io"
	"log"

	"github.com/simonwaldherr/golibs/ansi"
)

type Terminal struct {
	statusCode int
	path       []P
	mw         io.Writer
	*bytes.Buffer
	*log.Logger
}

func NewTerminal(w io.Writer) *Terminal {
	buff := bytes.NewBuffer(nil)
	mw := io.MultiWriter(w, buff)

	return &Terminal{
		mw:     mw,
		Buffer: buff,
		Logger: log.New(mw, "", 0),
	}
}

func (t *Terminal) prefix() string {
	fix := ""
	for _, _ = range t.path {
		fix += "   "
	}
	return fix
}

func (t *Terminal) Pipe() io.Writer {
	r, w := io.Pipe()
	s := bufio.NewScanner(r)
	go func() {
		for s.Scan() {
			t.Print(s.Text())
		}
	}()

	return w
}

func (t *Terminal) StatusCode() int {
	return t.statusCode
}

func (t *Terminal) SetStatusCode(code int) {
	t.statusCode = code
}

func (t *Terminal) Success(stepFn StepFunc, args ...interface{}) {
	_, str := stepFn(args...)
	t.Print(ansi.Color(str, ansi.Green))
}

func (t *Terminal) Warning(stepFn StepFunc, args ...interface{}) {
	_, str := stepFn(args...)
	t.Print(ansi.Color(str, ansi.Yellow))
}

func (t *Terminal) Error(stepFn StepFunc, args ...interface{}) {
	_, str := stepFn(args...)
	t.Print(ansi.Color(str, ansi.Red))
}

func (t *Terminal) Report(stepFn StepFunc, args ...interface{}) {
	_, str := stepFn(args...)
	t.Print(str)
}

func (t *Terminal) Enter(p P, stepFn StepFunc, args ...interface{}) {
	if stepFn != nil {
		_, str := stepFn(args...)
		t.Print(ansi.Underline(str))
	}

	t.path = append(t.path, p)
	t.Logger.SetPrefix(t.prefix())
}

func (t *Terminal) Exit() {
	t.path = t.path[0 : len(t.path)-1]
	t.Logger.SetPrefix(t.prefix())
}

func (t *Terminal) Path() string {
	path := ""
	for _, p := range t.path {
		path += "." + p.ID()
	}
	return path
}
