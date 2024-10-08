package p1497

func canArrange(arr []int, k int) bool {
	freq := map[int]int{}
	for _, v := range arr {
		e := ((v % k) + k) % k
		freq[e]++
	}
	if freq[0]%2 != 0 {
		return false
	}
	delete(freq, 0)
	for e, v := range freq {
		rem := k - e
		if rv, ok := freq[rem]; !ok {
			return false
		} else if rv != v {
			return false
		}
	}
	return true
}
