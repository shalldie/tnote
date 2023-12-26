package utils

// 三元运算
func Ternary[T any](condition bool, item1 T, item2 T) T {
	if condition {
		return item1
	}

	return item2
}

// 获取较大的数
func MathMax(a int, b int) int {
	return Ternary(a > b, a, b)
}
