package command

import (
	"fmt"
	"io"
	"text/template"

	"github.com/codegangsta/cli"
)

var tmpl_build = `State building successful!`

type Build struct {
	*cmd
}

func NewBuild(out io.Writer) *Build {
	return &Build{
		cmd: newCmd(out),
	}
}

func (c *Build) Name() string {
	return "build"
}

func (c *Build) Description() string {
	return fmt.Sprintf("...")
}

func (c *Build) Usage() string {
	return "Build the states using Docker"
}

func (c *Build) Flags() []cli.Flag {
	fs := []cli.Flag{}
	fs = append(fs, c.ConfigFlags()...)
	fs = append(fs, c.ParseExampleFlags()...)
	fs = append(fs, c.BuildStatesFlags()...)
	fs = append(fs, c.DockerFlags()...)

	return fs
}

func (c *Build) Action() func(ctx *cli.Context) {
	return c.templated(c.Run)
}

func (c *Build) Run(ctx *cli.Context) (*template.Template, interface{}, error) {

	//get manifest
	m, err := c.ParseExamples(ctx)
	if err != nil {
		return nil, nil, err
	}

	//get all states in the manifest
	states, err := m.States()
	if err != nil {
		return nil, nil, err
	}

	//load configuration
	conf, err := c.LoadConfig(ctx)
	if err != nil {
		return nil, nil, err
	}

	//add default states
	for _, pc := range conf.ProviderConfigs() {
		if sts, ok := states[pc.Name()]; ok {
			states[pc.Name()] = append(sts, pc.DefaultState())
		}
	}

	//get the state manager
	sm, err := c.StateManager(ctx)
	if err != nil {
		return nil, nil, err
	}

	//loop over states and build them
	for pname, snames := range states {
		for _, sname := range snames {
			fmt.Fprintf(c.out, "Building state %s/%s:\n", pname, sname)
			iname, err := sm.Build(pname, sname, c.out)
			if err != nil {
				fmt.Fprintf(c.out, "ERROR \n")
				return nil, nil, err
			}

			fmt.Fprintf(c.out, "done! (%s)\n\n", iname)
		}
	}

	return template.Must(template.New("build.success").Parse(tmpl_build)), nil, nil
}
