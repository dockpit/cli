package command_test

import (
	"bytes"
	"testing"

	"github.com/dockpit/pit/command"

	"github.com/dockpit/pit/spec"
)

type mock1 struct{}

func (m *mock1) StopAll() error                 { return nil }
func (m *mock1) Start(*spec.Dependencies) error { return nil }

func TestSingleSpecMocking(t *testing.T) {

	out := bytes.NewBuffer(nil)
	cmd := command.NewMock(out)

	//mock docker client
	cmd.Docker = &mock1{}

	//@todo implement check
	AssertCommand(t, cmd, []string{"-d=tcp://bogus", "../examples/notes.json"}, `(?s).*`, out)
}
