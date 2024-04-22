package proenv

import (
	"errors"
	"os"
	"reflect"
)

func LoadEnv[T interface{}](finalEnv *T) error {
	if finalEnv == nil {
		return errors.New("[proenv]: finalEnv cannot be a nil pointer")
	}
	envType := reflect.TypeOf(*finalEnv)
	envEditable := reflect.ValueOf(finalEnv).Elem()
	for i := 0; i < envType.NumField(); i++ {
		field := envType.Field(i)
		if field.Type.Kind() != reflect.String {
			return errors.New("[proenv]: All Struct Env fields must be of type string")
		}
		envTag := field.Tag.Get("env")
		if envTag == "" {
			return errors.New("[proenv]: All Struct Env fields must have a env tag")
		}
		fieldValue := envEditable.FieldByName(field.Name)
		if fieldValue.IsValid() && fieldValue.CanSet() {
			fieldValue.SetString(os.Getenv(envTag))
		}
	}
	return nil
}
