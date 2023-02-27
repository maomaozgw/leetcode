package p427

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	return constructPart(grid, 0, 0, len(grid)-1, len(grid)-1)
}

func constructPart(grid [][]int, leftX, leftY, rightX, rightY int) *Node {
	if val, isLeaf := IsLeaf(grid, leftX, leftY, rightX, rightY); isLeaf {
		return &Node{
			Val:    val,
			IsLeaf: isLeaf,
		}
	}
	node := &Node{
		Val:    false,
		IsLeaf: false,
	}
	centerX := (leftX + rightX) / 2
	centerY := (leftY + rightY) / 2
	node.TopLeft = constructPart(grid, leftX, leftY, centerX, centerY)
	node.TopRight = constructPart(grid, leftX, centerY+1, centerX, rightY)
	node.BottomLeft = constructPart(grid, centerX+1, leftY, rightX, centerY)
	node.BottomRight = constructPart(grid, centerX+1, centerY+1, rightX, rightY)
	return node
}

func IsLeaf(grid [][]int, leftX, leftY, rightX, rightY int) (bool, bool) {
	firstVal := grid[leftX][leftY]
	if rightX == leftX {
		return firstVal == 1, true
	}
	for i := leftX; i <= rightX; i++ {
		for j := leftY; j <= rightY; j++ {
			if grid[i][j] != firstVal {
				return false, false
			}
		}
	}
	return firstVal == 1, true
}
