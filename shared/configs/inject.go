package configs

import (
	"reflect"

	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

var registeredInjects = make(map[string]interface{})

func RegisterInject[T interface{}](key string, value *T) {
	registeredInjects[key] = value
}

func ResolveInject[T interface{}](instance *T) *T {
	envType := reflect.TypeOf(*instance)
	envEditable := reflect.ValueOf(instance).Elem()

	for i := 0; i < envType.NumField(); i++ {
		field := envType.Field(i)
		tag := field.Tag.Get("inject")
		if tag == "" {
			continue
		}
		invalidTag := true
		for key := range registeredInjects {
			if tag == key {
				invalidTag = false
				break
			}
		}
		utils.Assert(!invalidTag, "[Configs]: Invalid `inject` value")
		fieldValue := envEditable.FieldByName(field.Name)
		if fieldValue.IsValid() && fieldValue.CanSet() {
			fieldValue.Set(reflect.ValueOf(registeredInjects[tag]))
		}
	}
	return instance
}
