package p2924

func findChampion(n int, edges [][]int) int {
	var weakers = make([]bool, n)
	for _, v := range edges {
		weakers[v[1]] = true
	}
	var stronger int = -1
	for i := 0; i < n; i++ {
		if !weakers[i] {
			if stronger != -1 {
				return -1
			}
			stronger = i
		}
	}
	return stronger
}

func findChampionWithMap(n int, edges [][]int) int {
	var strongers = map[int]bool{}
	for i := 0; i < n; i++ {
		strongers[i] = true
	}
	for _, v := range edges {
		delete(strongers, v[1])
	}

	if len(strongers) == 1 {
		for v := range strongers {
			return v
		}
	}
	return -1
}
