package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type LogConfig struct {
	Filename string    `mapstructure:"filename"`
	Level    log.Level `mapstructure:"level"`
}

type Config struct {
	Log LogConfig `mapstructure:"log"`
}

func Parse(v *viper.Viper) (Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg, viper.DecodeHook(LogLevelHook))
	return cfg, err
}
