package redditutils

const (
	MAX_LIMIT_PER_REQUEST = 100
)

func RequestLimit(current int, total int) int {
	d := total - current
	if d < MAX_LIMIT_PER_REQUEST {
		return d
	} else {
		return MAX_LIMIT_PER_REQUEST
	}
}
