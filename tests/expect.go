package tests

import "fmt"

func ExpectString[T interface{}](actual T, expected T) string {
	return fmt.Sprintf("Expected: %v\nGot: %v", expected, actual)
}
