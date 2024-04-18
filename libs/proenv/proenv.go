package proenv

import (
	"errors"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

func LoadEnv[T interface{}](finalEnv *T) error {
	err := godotenv.Load()
	if err != nil {
		return errors.New("[proenv]: You must have a .env file for the env variables to be loaded")
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
