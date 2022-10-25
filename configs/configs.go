package configs

import (
	"github.com/kelseyhightower/envconfig"
)

var Settings Config

type Config struct {
	ServerPort int `required:"true" split_words:"true" default:"8080"`
}

func Load() error {
	err := envconfig.Process("app", &Settings)
	if err != nil {
		return err
	}

	return nil
}
