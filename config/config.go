package config

import (
	"gopkg.in/yaml.v2"
)

type ApiKeys struct {
	Telegram string
}

type Config struct {
	ApiKeys ApiKeys
}

type configFile struct {
	APIKeys struct {
		Telegram string `yaml:"telegram"`
	} `yaml:"apiKeys"`
}

func ParseConfig(fileBytes []byte) (*Config, error) {
	cf := configFile{}
	err := yaml.Unmarshal(fileBytes, &cf)
	if err != nil {
		return nil, err
	}

	c := Config{}
	c.ApiKeys.Telegram = cf.APIKeys.Telegram

	return &c, nil
}
