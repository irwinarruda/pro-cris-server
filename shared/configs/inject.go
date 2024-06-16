package configs

import (
	"reflect"

	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

func ResolveInject[T interface{}](instance *T) *T {
	env := GetEnv()
	validate := GetValidate()
	db := GetDb()

	envType := reflect.TypeOf(*instance)
	envEditable := reflect.ValueOf(instance).Elem()

	for i := 0; i < envType.NumField(); i++ {
		field := envType.Field(i)
		tag := field.Tag.Get("inject")
		if tag == "" {
			continue
		}
		utils.Assert(tag == "env" || tag == "validate" || tag == "db", "[Configs]: Invalid `inject` value")
		fieldValue := envEditable.FieldByName(field.Name)
		if fieldValue.IsValid() && fieldValue.CanSet() {
			if tag == "env" {
				fieldValue.Set(reflect.ValueOf(env))
			}
			if tag == "validate" {
				fieldValue.Set(reflect.ValueOf(validate))
			}
			if tag == "db" {
				fieldValue.Set(reflect.ValueOf(db))
			}
		}
	}
	return instance
}
