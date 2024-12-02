package utils

func Diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
