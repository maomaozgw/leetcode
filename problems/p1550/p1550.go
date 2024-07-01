package p1550

func threeConsecutiveOdds(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	for i := 0; i < len(arr)-2; i++ {
		if arr[i]%2 == 1 && arr[i+1]%2 == 1 && arr[i+2]%2 == 1 {
			return true
		}
	}
	return false
}
