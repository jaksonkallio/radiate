package config

import (
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var CurrentConfig Config

type Config struct {
	MinLogLevel     zerolog.Level `yaml:"min_log_level"`
	CacheDir        string        `yaml:"cache_dir"`
	IPFSHost        string        `yaml:"ipfs_host"`
	APIHost         string        `yaml:"api_host"`
	PrettyPrintLogs bool          `yaml:"pretty_print_logs"`
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
