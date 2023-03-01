package p860

func lemonadeChange(bills []int) bool {
	var change = make([]int, 3)
	for i := 0; i < len(bills); i++ {
		if !doChange(change, bills[i]) {
			return false
		}
	}
	return true
}

func doChange(changes []int, bill int) bool {
	switch bill {
	case 5:
		changes[0]++
	case 10:
		if changes[0] == 0 {
			return false
		}
		changes[1]++
		changes[0]--
	case 20:
		if changes[1] == 0 {
			if changes[0] < 3 {
				return false
			}
			changes[0] -= 3
			break
		}
		if changes[0] == 0 {
			return false
		}
		changes[1]--
		changes[0]--
	default:
		panic("bill issue")
	}
	return true
}
