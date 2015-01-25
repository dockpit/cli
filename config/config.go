package config

import (
	"fmt"
	"regexp"
	"time"
)

type Config struct {
	data *ConfigData

	depConfigs []*DependencyConfig
	spConfigs  []*StateProviderConfig
	runConfig  *RunConfig
	subject    string
}

func Parse(cd *ConfigData) (*Config, error) {

	//parse deps into port configs
	depsconf := []*DependencyConfig{}
	for dep, confs := range cd.Dependencies {
		ports := []*PortConfig{}
		for _, pdata := range *confs {
			pconf, err := ParsePort(pdata)
			if err != nil {
				return nil, err
			}

			ports = append(ports, pconf)
		}

		depsconf = append(depsconf, &DependencyConfig{
			Name:  dep,
			Ports: ports,
		})
	}

	//parse state provider config
	spconf := []*StateProviderConfig{}
	for pname, conf := range cd.StateProviders {
		ports := []*PortConfig{}
		for _, pdata := range conf.Ports {
			pconf, err := ParsePort(pdata)
			if err != nil {
				return nil, err
			}

			ports = append(ports, pconf)
		}

		//get regexp to determine when the state provider is ready
		exp, err := regexp.Compile(conf.ReadyPattern)
		if err != nil {
			return nil, err
		}

		//get timeout by parsing duration
		if conf.ReadyTimeout == "" {
			conf.ReadyTimeout = "2s"
		}

		d, err := time.ParseDuration(conf.ReadyTimeout)
		if err != nil {
			return nil, err
		}

		//default state
		if conf.DefaultState == "" {
			conf.DefaultState = "default"
		}

		spconf = append(spconf, &StateProviderConfig{
			name:         pname,
			ports:        ports,
			readyExp:     exp,
			cmd:          conf.Cmd,
			readyTimeout: d,
			defaultState: conf.DefaultState,
		})
	}

	if cd.Run == nil {
		cd.Run = &RunData{}
	}

	rconf, err := ParseRunConfig(cd.Run)
	if err != nil {
		return nil, err
	}

	if cd.Subject == "" {
		return nil, fmt.Errorf("No test subject specified in configuration file")
	}

	return &Config{cd, depsconf, spconf, rconf, cd.Subject}, nil
}

func (c *Config) Subject() string {
	return c.subject
}

func (c *Config) Data() *ConfigData {
	return c.data
}

func (c *Config) RunConfig() *RunConfig {
	return c.runConfig
}

func (c *Config) StateProviderConfig(pname string) StateProviderC {
	for _, spc := range c.spConfigs {
		if spc.name == pname {
			return spc
		}
	}

	return nil
}

func (c *Config) PortsForStateProvider(pname string) []*PortConfig {
	for _, spc := range c.spConfigs {
		if spc.name == pname {
			return spc.Ports()
		}
	}

	return nil
}

func (c *Config) PortsForDependency(dep string) []*PortConfig {
	for _, depc := range c.depConfigs {
		if depc.Name == dep {
			return depc.Ports
		}
	}

	return nil
}

func (c *Config) DependencyConfigs() []DependencyC {
	res := []DependencyC{}

	for _, depc := range c.depConfigs {
		res = append(res, depc)
	}

	return res
}

func (c *Config) ProviderConfigs() []StateProviderC {
	res := []StateProviderC{}

	for _, spc := range c.spConfigs {
		res = append(res, spc)
	}

	return res
}
