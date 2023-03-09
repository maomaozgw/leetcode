package p2073

func timeRequiredToBuy(tickets []int, k int) int {
	var result = 0
	var max = tickets[k]
	for i := range tickets {
		currentTime := tickets[i]
		diff := currentTime - max
		if i < k {
			result += max
			if diff < 0 {
				result += diff
			}
		} else if i > k {
			if diff >= 0 {
				result += max - 1
			} else {
				result += currentTime
			}
		}
	}
	return result + max
}
