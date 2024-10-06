package utils

import (
	"errors"
	"slices"
	"strings"
)

func UnmarshalEnum[T ~string](enum *T, enumValues []string, b []byte) error {
	value := strings.Trim(string(b), "\"")
	if !slices.Contains(enumValues, value) {
		return errors.New("invalid enum value")
	}
	*enum = T(value)
	return nil
}
