package config

import (
	"reflect"

	log "github.com/sirupsen/logrus"
)

func LogLevelHook(f, t reflect.Type, data any) (any, error) {
	if f.Kind() == reflect.String && t == reflect.TypeOf(log.PanicLevel) {
		return log.ParseLevel(data.(string))
	}
	return data, nil
}
