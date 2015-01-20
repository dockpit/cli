package config

type DependencyConfigData []string

type StateProviderConfigData struct {
	Ports        []string `json:"ports"`
	ReadyPattern string   `json:"ready_pattern"`
	ReadyTimeout string   `json:"ready_timeout"`
	DefaultState string   `json:"default_state"`
	Cmd          []string `json:"command"`
}

type RunData struct {
	Command      []string `json:"command"`
	Dir          string   `json:"dir"`
	ReadyPattern string   `json:"ready_pattern"`
	ReadyTimeout string   `json:"ready_timeout"`
}

type ConfigData struct {
	Dependencies   map[string]*DependencyConfigData    `json:"deps"`
	StateProviders map[string]*StateProviderConfigData `json:"states"`
	Run            *RunData                            `json:"run"`
	DockerHostname string                              `json:"docker_hostname"`
	Subject        string                              `json:"subject"`
}

var InitConfigData = `{
	"subject": "http://localhost",
	"run": {
		"command": ["start", "your", "server"]
	}
}`
