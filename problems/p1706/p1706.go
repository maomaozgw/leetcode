// https://leetcode.com/problems/where-will-the-ball-fall/

package p1706

const (
	fromTop   = 0
	fromLeft  = 1
	fromRight = 2
)

type pos struct {
	x, y int
}

type cacheInfo struct {
	prev pos
	cur  pos
}

func findBall(grid [][]int) []int {
	var n = len(grid)
	if n == 0 {
		return []int{}
	}
	var l = len(grid[0])
	var result = make([]int, l)
	var cache = map[cacheInfo]int{}
	for i := 0; i < l; i++ {
		result[i] = findResult(grid, pos{x: -1, y: i}, pos{x: 0, y: i}, cache)
	}
	return result
}

func findResult(grid [][]int, prev, cur pos, cache map[cacheInfo]int) int {
	if val, ok := cache[cacheInfo{
		prev: prev,
		cur:  cur,
	}]; ok {
		return val
	}
	var next = cur
	val := grid[cur.x][cur.y]
	var from = fromTop
	if prev.y == cur.y {
		from = fromTop
	} else if prev.y > cur.y {
		from = fromRight
	} else {
		from = fromLeft
	}
	switch from {
	case fromTop:
		next.y += val
	case fromRight:
		if val == 1 {
			next.y++
		} else {
			next.x++
		}
	case fromLeft:
		if val == -1 {
			next.y--
		} else {
			next.x++
		}
	default:
		panic("cannot here")
	}
	if next.y < 0 || next.y >= len(grid[0]) {
		return -1
	}
	if next.x == len(grid) {
		return next.y
	}
	if next == prev {
		return -1
	}
	result := findResult(grid, cur, next, cache)
	cache[cacheInfo{
		prev: prev,
		cur:  cur,
	}] = result
	return result
}
