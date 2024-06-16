package utils

func Map[T interface{}, K interface{}](arr []T, fn func(T) K) []K {
	var result []K
	for _, v := range arr {
		result = append(result, fn(v))
	}
	return result
}
