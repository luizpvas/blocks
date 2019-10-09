package config

import (
	"fmt"

	"github.com/luizpvas/blocks/records"
	"gopkg.in/yaml.v2"
)

// AppConfig is describes the structure of a blocks app configuration file.
type AppConfig struct {
	HTTP      HTTP `yaml:"http"`
	Resources map[string]*records.Resource
}

// HTTP configuration for blocks server.
type HTTP struct {
	Listen         string
	StaticFilesDir string `yaml:"static_files_dir"`
}

// ParseAppConfig attempts to parse the configuration from the YAML file.
func ParseAppConfig(data []byte) (*AppConfig, error) {
	appconfig := &AppConfig{
		HTTP: HTTP{
			Listen: "127.0.0.1:8080",
		},
	}

	err := yaml.Unmarshal(data, appconfig)
	if err != nil {
		return nil, fmt.Errorf("could not parse config from YAML file: %v", err)
	}

	for key, resource := range appconfig.Resources {
		resource.ID = key
		for name, field := range resource.Fields {
			field.Name = name
		}
	}

	return appconfig, nil
}
