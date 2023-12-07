package config

import (
	"encoding/json"
	"os"

	"github.com/google/go-jsonnet"
)

type Config struct {
	ConfigPath   string   `json:"config_path,omitempty"`
	Dependencies []string `json:"deps,omitempty"`
	Setup        Steps    `json:"setup,omitempty"`
	Spec         string   `json:"spec,omitempty"`
}

func LoadConfigFile(path string) (*Config, error) {
	vm := jsonnet.MakeVM()
	jsonData, err := vm.EvaluateFile(path)
	if err != nil {
		return nil, err
	}
	pkg := &Config{}
	err = json.Unmarshal([]byte(jsonData), pkg)
	if err != nil {
		return nil, err
	}

	pkg.ConfigPath = path

	return pkg, nil
}

func Save(path string, cfg *Config) error {
	out, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, out, 0o755)
}
