package p2421

import "sort"

func numberOfGoodPaths(vals []int, edges [][]int) int {
	var uf = make([]int, len(vals))
	for i := 0; i < len(uf); i++ {
		uf[i] = i
	}
	var find = func(i int) (root int) {
		for uf[i] != i {
			i = uf[i]
		}
		return i
	}
	var union = func(i, j int) {
		r1, r2 := find(i), find(j)
		uf[r1] = r2
	}
	var tree = make([][]int, len(vals))
	for _, val := range edges {
		start, end := val[0], val[1]
		tree[start] = append(tree[start], end)
		tree[end] = append(tree[end], start)
	}
	var valuesMap = map[int][]int{}
	var keys []int
	for i := 0; i < len(vals); i++ {
		if _, ok := valuesMap[vals[i]]; !ok {
			keys = append(keys, vals[i])
		}
		valuesMap[vals[i]] = append(valuesMap[vals[i]], i)
	}

	sort.Ints(keys)

	var goodPath = 0

	for _, val := range keys {
		nodes := valuesMap[val]
		for _, node := range nodes {
			for _, neighbor := range tree[node] {
				if vals[node] >= vals[neighbor] {
					union(node, neighbor)
				}
			}
		}
		group := map[int]int{}
		for _, node := range nodes {
			group[find(node)]++
		}
		for _, size := range group {
			goodPath += size * (size + 1) / 2
		}
	}
	return goodPath
}
