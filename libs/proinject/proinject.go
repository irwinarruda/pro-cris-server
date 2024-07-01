package proinject

import (
	"log"
	"reflect"
)

var registeredInjects = make(map[string]interface{})

func Register[T interface{}](key string, value *T) {
	registeredInjects[key] = value
}

func Resolve[T interface{}](instance *T) *T {
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
		if invalidTag {
			log.Fatal("[proinject]: Invalid `inject` value in struct field. ", tag)
		}
		fieldValue := envEditable.FieldByName(field.Name)
		if fieldValue.IsValid() && fieldValue.CanSet() {
			fieldValue.Set(reflect.ValueOf(registeredInjects[tag]))
		}
	}
	return instance
}
