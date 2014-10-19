package command_test

import (
	"os"
	"testing"

	"github.com/dockpit/pit/command"
	"github.com/dockpit/pit/spec"
)

func getclient(t *testing.T) command.D {
	host := os.Getenv("DOCKER_HOST")
	if host == "" {
		t.Skip("No DOCKER_HOST env variable setup")
	}

	d, err := command.NewDocker(host)
	if err != nil {
		t.Fatal(err)
	}

	return d
}

func TestStart(t *testing.T) {
	var d command.D
	d = getclient(t)
	_ = d

	//get speck
	l := spec.NewLoader()
	f := spec.NewFactory(l)
	s, err := f.Create("../examples/notes.json")
	if err != nil {
		t.Fatal(err)
	}

	deps, err := s.Dependencies()
	if err != nil {
		t.Fatal(err)
	}

	err = d.Start(deps)
	if err != nil {
		t.Fatal(err)
	}

}
