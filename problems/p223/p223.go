// https://leetcode.com/problems/rectangle-area/

package p223

import "fmt"

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	aa := calcRectangleArea(ax1, ay1, ax2, ay2)
	ba := calcRectangleArea(bx1, by1, bx2, by2)
	var leftX1, leftX2, rightX1, rightX2 = ax1, ax2, bx1, bx2
	var bottomY1, bottomY2, topY1, topY2 = ay1, ay2, by1, by2
	if leftX1 > rightX1 {
		leftX1, rightX1 = bx1, ax1
	}
	if bottomY1 > topY1 {
		bottomY1, topY1 = by1, ay1
	}
	if leftX2 > rightX2 {
		leftX2, rightX2 = bx2, ax2
	}
	if bottomY2 > topY2 {
		bottomY2, topY2 = by2, ay2
	}
	fmt.Println(aa, ba)
	if rightX1 >= leftX2 || topY1 >= bottomY2 {
		// no overlap
		return aa + ba
	}

	ca := calcRectangleArea(rightX1, topY1, leftX2, bottomY2)

	return aa + ba - ca
}

func calcRectangleArea(x1, y1, x2, y2 int) int {
	return (x2 - x1) * (y2 - y1)
}
