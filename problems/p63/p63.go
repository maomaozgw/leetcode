// https://leetcode.com/problems/unique-paths-ii/

package p63

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	} else if obstacleGrid[0][0] == 1 {
		return 0
	}
	if m == 1 && n == 1 {
		return 1
	}

	var rMap [][]int
	rMap = make([][]int, m)
	for i := 0; i < m; i++ {
		rMap[i] = make([]int, n)
	}
	return walkInto(obstacleGrid, rMap, 0, 0, m-1, n-1)
}

func walkInto(robotMap [][]int, knownMap [][]int, currentX, currentY, endX, endY int) (ways int) {
	if val := knownMap[currentX][currentY]; val != 0 {
		return val
	}

	nextY := currentY + 1
	if currentX == endX && nextY == endY {
		return 1
	}
	if nextY <= endY && robotMap[currentX][nextY] == 0 {
		knownMap[currentX][currentY] += walkInto(robotMap, knownMap, currentX, nextY, endX, endY)
	}

	nextX := currentX + 1
	if currentY == endY && nextX == endX {
		return 1
	}
	if nextX <= endX && robotMap[nextX][currentY] == 0 {
		knownMap[currentX][currentY] += walkInto(robotMap, knownMap, nextX, currentY, endX, endY)
	}
	return knownMap[currentX][currentY]
}
