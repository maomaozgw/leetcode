package p733

var (
	xNext = []int{1, -1, 0, 0}
	yNext = []int{0, 0, 1, -1}
)

const (
	offset = 100
)

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if len(image) == 0 {
		return image
	}
	var startPoint = sr*offset + sc
	var q = []int{startPoint}
	var cVal = image[sr][sc]
	if cVal == color {
		return image
	}
	image[sr][sc] = color
	var m, n = len(image), len(image[0])
	for len(q) > 0 {
		var newQ []int
		for j := 0; j < len(q); j++ {
			x, y := q[j]/offset, q[j]%offset
			for i := 0; i < 4; i++ {
				nX := x + xNext[i]
				nY := y + yNext[i]
				if nX < 0 || nY < 0 || nX == m || nY == n {
					continue
				}
				if image[nX][nY] == cVal {
					image[nX][nY] = color
					newQ = append(newQ, nX*offset+nY)
				}
			}
		}
		q = newQ
	}
	return image
}
