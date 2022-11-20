// https://leetcode.com/problems/erect-the-fence/

package p587

import "sort"

func outerTrees(trees [][]int) [][]int {
	sortT(trees)
	trees = build(trees)

	var out [][]int
	var set = map[[2]int]bool{}
	for _, t := range trees {
		if k := [2]int{t[0], t[1]}; !set[k] {
			out = append(out, t)
			set[k] = true
		}
	}
	return out
}

func build(points [][]int) [][]int {
	cross := func(o, a, b []int) int {
		return (a[0]-o[0])*(b[1]-o[1]) - (a[1]-o[1])*(b[0]-o[0])
	}
	var t, b [][]int
	for _, p := range points {
		for len(t) >= 2 && cross(t[len(t)-2], t[len(t)-1], p) > 0 {
			t = t[:len(t)-1]
		}
		for len(b) >= 2 && cross(b[len(b)-2], b[len(b)-1], p) < 0 {
			b = b[:len(b)-1]
		}
		t, b = append(t, p), append(b, p)
	}
	return append(t, b...)
}

func sortT(trees [][]int) {
	// 优先按照X坐标升序，如果X坐标一致按照Y升序
	sort.Slice(trees, func(i, j int) bool {
		if trees[i][0] != trees[j][0] {
			return trees[i][0] < trees[j][0]
		}
		return trees[i][1] < trees[j][1]
	})
}
