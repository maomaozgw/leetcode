package p1011

func shipWithinDays(weights []int, days int) int {
	var canShip = func(cap int) bool {
		s, d := 0, 0

		for i := 0; i < len(weights); i++ {
			w := weights[i]
			if s+w <= cap {
				s += w
			} else {
				d++
				s = w
			}
			if d >= days {
				break
			}
		}
		return d < days
	}
	var left, right = 0, 0
	for idx := range weights {
		if weights[idx] > left {
			left = weights[idx]
		}
		right += weights[idx]
	}
	if right/days > left {
		left = right / days
	}
	for left < right {
		current := (right + left) / 2
		if !canShip(current) {
			left = current + 1
		} else {
			right = current
		}
	}
	return left
}
