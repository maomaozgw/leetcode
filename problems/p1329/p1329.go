// https: //leetcode.com/problems/sort-the-matrix-diagonally/

package p1329

import (
	"github.com/emirpasic/gods/trees/binaryheap"
)

func diagonalSort(mat [][]int) [][]int {
	rl := len(mat)
	cl := len(mat[0])
	h := binaryheap.NewWithIntComparator()
	for i := 0; i < cl; i++ {
		for j := 0; j < rl; j++ {
			key := i + j
			if key == cl {
				break
			}
			h.Push(mat[j][key])
		}
		for j := 0; j < rl; j++ {
			key := i + j
			if key == cl {
				break
			}
			val, _ := h.Pop()
			mat[j][key] = val.(int)
		}
	}

	for i := 1; i < rl; i++ {
		for j := 0; j < cl; j++ {
			key := i + j
			if key == rl {
				break
			}
			h.Push(mat[key][j])
		}

		for j := 0; j < cl; j++ {
			key := i + j
			if key == rl {
				break
			}
			val, _ := h.Pop()
			mat[key][j] = val.(int)
		}
	}
	return mat
}
