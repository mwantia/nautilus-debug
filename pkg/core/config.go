package core

import "github.com/mitchellh/mapstructure"

type DebugConfig struct {
	Name string `mapstructure:"name,omitempty"`
}

func GetConfig(data map[string]interface{}) (*DebugConfig, error) {
	var cfg *DebugConfig
	if err := mapstructure.Decode(data, &cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
