package config

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

type Config struct {
	DbUser     string `yaml:"db_user"`
	DbPassword string `yaml:"db_password"`
	DbHost     string `yaml:"db_host"`
	DbPort     uint16 `yaml:"db_port"`
	DbName     string `yaml:"db_name"`
}

func ParseConfig(configPath string) (*Config, error) {
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
