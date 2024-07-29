package p1395

func numTeams(rating []int) int {
	var result int
	for i := 1; i < len(rating)-1; i++ {
		var leftLess, leftGreater, rightLess, rightGreater int
		for j := 0; j < i; j++ {
			if rating[j] < rating[i] {
				leftLess++
			} else if rating[j] > rating[i] {
				leftGreater++
			}
		}
		for j := i + 1; j < len(rating); j++ {
			if rating[j] < rating[i] {
				rightLess++
			} else if rating[j] > rating[i] {
				rightGreater++
			}
		}
		result += leftLess*rightGreater + leftGreater*rightLess
	}
	return result
}
