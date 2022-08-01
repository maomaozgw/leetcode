// https://leetcode.com/problems/unique-paths/

package p62

func uniquePaths(m int, n int) int {
	if m == 1 && n == 1 {
		return 1
	} else if m == 1 {
		return 1
	} else if n == 1 {
		return 1
	}
	var rMap [][]int
	rMap = make([][]int, m)
	for i := 0; i < m; i++ {
		rMap[i] = make([]int, n)
	}
	return walkInto(rMap, 0, 0, m-1, n-1)
}

func walkInto(knownMap [][]int, currentX, currentY, endX, endY int) (ways int) {
	if val := knownMap[currentX][currentY]; val != 0 {
		return val
	}

	nextY := currentY + 1
	if currentX == endX && nextY == endY {
		return 1
	}
	if nextY <= endY {
		knownMap[currentX][currentY] += walkInto(knownMap, currentX, nextY, endX, endY)
	}

	nextX := currentX + 1
	if currentY == endY && nextX == endX {
		return 1
	}
	if nextX <= endX {
		knownMap[currentX][currentY] += walkInto(knownMap, nextX, currentY, endX, endY)
	}
	return knownMap[currentX][currentY]
}
