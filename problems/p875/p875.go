package p875

func minEatingSpeed(piles []int, h int) int {
	var start, end = 1, 10000000000
	for start <= end {
		var trip = 0
		var mid = (start + end) / 2
		for i := 0; i < len(piles); i++ {
			trip += piles[i] / mid
			if piles[i]%mid > 0 {
				trip++
			}
		}
		if trip > h {
			start = mid + 1
			continue
		}
		end = mid - 1
	}
	return start
}
