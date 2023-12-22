package utils

func Ternary[T any](condition bool, item1 T, item2 T) T {
	if condition {
		return item1
	}

	return item2
}
