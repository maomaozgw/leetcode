package p1903

const (
	start = byte('0')
)

func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		current := num[i] - start
		if current%2 == 0 {
			continue
		}
		return num[:i+1]
	}
	return ""
}
