package config

import (
	"fennec/core"

	configUtil "github.com/fox-one/pkg/config"
)

// Load load config file
func Load(configFile string, config *core.Config) error {
	configUtil.AutomaticLoadEnv("FENNEC")
	if err := configUtil.LoadYaml(configFile, config); err != nil {
		return err
	}

	return nil
}
