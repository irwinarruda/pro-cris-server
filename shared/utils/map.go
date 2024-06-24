package utils

func Map[T interface{}, K interface{}](arr []T, fn func(T, int) K) []K {
	var result = []K{}
	for i, v := range arr {
		result = append(result, fn(v, i))
	}
	return result
}
