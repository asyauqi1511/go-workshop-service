package config

import (
	"os"

	"github.com/asyauqi1511/go-workshop-service/internal/entity"
	"github.com/asyauqi1511/go-workshop-service/internal/pkg/errorwrapper"
	"gopkg.in/yaml.v3"
)

func LoadConfig() (entity.AppConfig, error) {

	var config entity.AppConfig

	// Read YAML file.
	yamlFile, err := os.ReadFile("files/config/config.development.yaml")
	if err != nil {
		return config, errorwrapper.Wrap(err)
	}

	// Unmarhal YAML to config.
	err = yaml.Unmarshal(yamlFile, &config)

	return config, errorwrapper.Wrap(err)
}
