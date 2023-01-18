package utils

func RequestLimit(current int, total int) int {
	d := total - current
	if d < 100 {
		return d
	} else {
		return 100
	}
}
