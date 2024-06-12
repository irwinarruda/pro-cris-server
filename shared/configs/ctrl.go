package configs

import (
	"reflect"

	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

func ResolveCtrl[T interface{}](ctrl *T) *T {
	env := GetEnv()
	validate := GetValidate()
	db := GetDb()

	envType := reflect.TypeOf(*ctrl)
	envEditable := reflect.ValueOf(ctrl).Elem()

	for i := 0; i < envType.NumField(); i++ {
		field := envType.Field(i)
		tag := field.Tag.Get("ctrl")
		utils.Assert(tag != "", "[Configs]: All Controller properties must have a `ctrl` tag")
		utils.Assert(tag == "env" || tag == "validate" || tag == "db", "[Configs]: Invalid `ctrl` value")
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
	return ctrl
}
