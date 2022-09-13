// https://leetcode.com/problems/the-number-of-weak-characters-in-the-game/

package p1996

const (
	maxNum = 100000
)

func numberOfWeakCharacters(properties [][]int) int {
	var dmap = make([]int, maxNum+1)
	for i := 0; i < len(properties); i++ {
		dmap[properties[i][0]] = max(dmap[properties[i][0]], properties[i][1])
	}
	for i := maxNum - 1; i >= 0; i-- {
		dmap[i] = max(dmap[i], dmap[i+1])
	}

	var count = 0
	for i := 0; i < len(properties); i++ {
		if properties[i][0] < maxNum && properties[i][1] < dmap[properties[i][0]+1] {
			count++
		}
	}
	return count
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
