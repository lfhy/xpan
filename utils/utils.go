package utils

func GetOneOrDefault[T any](v ...T) T {
	if len(v) > 0 {
		return v[0]
	}
	var t T
	return t
}
