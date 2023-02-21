package p815

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	var stations2Bus = map[int][]int{}
	for i := range routes {
		for _, station := range routes[i] {
			if _, ok := stations2Bus[station]; !ok {
				stations2Bus[station] = []int{}
			}
			stations2Bus[station] = append(stations2Bus[station], i)
		}
	}
	var q []int

	var lastBus = map[int]bool{}
	if buses, ok := stations2Bus[target]; !ok {
		return -1
	} else {
		for _, bus := range buses {
			lastBus[bus] = true
		}
	}
	if buses, ok := stations2Bus[source]; !ok {
		return -1
	} else {
		for _, bus := range buses {
			if lastBus[bus] {
				return 1
			}
			q = append(q, bus)
		}
	}
	var usedBus = map[int]bool{}
	var result = 0
	for len(q) > 0 {
		result++
		var newQ []int
		for _, bus := range q {
			usedBus[bus] = true
			for _, station := range routes[bus] {
				for _, bus := range stations2Bus[station] {
					if lastBus[bus] {
						return result + 1
					}
					if usedBus[bus] {
						continue
					}
					newQ = append(newQ, bus)
				}
			}
		}
		q = newQ
	}
	return -1
}
