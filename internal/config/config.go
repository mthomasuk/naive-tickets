package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Postgresql struct {
		Conn string `yaml:"conn"`
	} `yaml:"postgresql"`
	Stripe struct {
		BaseURL string `yaml:"baseurl"`
		Key     string `yaml:"key"`
	} `yaml:"stripe"`
}

// NewConfigFromFile retrurns config struct from a file path.
func NewFromFile(filepath string, cfg *Config) error {
	var cfgraw []byte
	var err error

	cfgraw, err = ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(cfgraw, &cfg); err != nil {
		return err
	}
	return nil
}
