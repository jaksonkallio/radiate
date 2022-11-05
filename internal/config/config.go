package config

import (
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var CurrentConfig Config

type Environment string

const (
	EnvironmentDev Environment = "dev"
	EnvironmentPrd Environment = "prd"
)

type Config struct {
	MinLogLevel zerolog.Level
	Environment Environment
}

func LoadFromFile() error {
	// TODO: make this configurable
	fileContent, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return err
	}

	var config Config

	err = yaml.Unmarshal(fileContent, &config)
	if err != nil {
		return err
	}

	CurrentConfig = config

	return nil
}

func (config Config) IsDev() bool {
	return config.Environment == EnvironmentDev
}
