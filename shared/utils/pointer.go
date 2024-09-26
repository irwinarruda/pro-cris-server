package utils

func ToP[T interface{}](v T) *T {
	return &v
}
