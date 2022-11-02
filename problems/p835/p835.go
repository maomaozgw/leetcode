// https://leetcode.com/problems/image-overlap/

package p835

type pos struct {
	x int
	y int
}

func largestOverlap(img1 [][]int, img2 [][]int) int {
	var s1 []pos
	var s2 []pos
	for i := 0; i < len(img1); i++ {
		for j := 0; j < len(img1[i]); j++ {
			if img1[i][j] == 1 {
				s1 = append(s1, pos{x: i, y: j})
			}
			if img2[i][j] == 1 {
				s2 = append(s2, pos{
					x: i,
					y: j,
				})
			}
		}
	}
	var cache = map[pos]int{}
	var result = 0
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			p := pos{
				x: s1[i].x - s2[j].x,
				y: s1[i].y - s2[j].y,
			}
			cache[p]++
			result = max(result, cache[p])
		}
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
