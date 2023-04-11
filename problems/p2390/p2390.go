package p2390

const (
	start = '*'
)

func removeStars(s string) string {
	var left = 0
	var result = make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == start {
			left--
			continue
		}
		result[left] = s[i]
		left++
	}
	return string(result[:left])
}
