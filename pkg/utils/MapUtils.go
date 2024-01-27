package utils

func Map[T, U any](f func(T) U, list []T) []U {
	result := make([]U, len(list))
	for i, v := range list {
		result[i] = f(v)
	}
	return result
}
