package config

import (
	"reflect"

	log "github.com/sirupsen/logrus"
)

func LogLevelHook(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if f.Kind() == reflect.String && t == reflect.TypeOf(log.PanicLevel) {
		return log.ParseLevel(data.(string))
	}
	return data, nil
}
