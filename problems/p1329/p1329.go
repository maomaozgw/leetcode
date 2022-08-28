// https: //leetcode.com/problems/sort-the-matrix-diagonally/

package p1329

import "sort"

func diagonalSort(mat [][]int) [][]int {
	rl := len(mat)
	cl := len(mat[0])
	for i := 0; i < cl; i++ {
		var line []int
		for j := 0; j < rl; j++ {
			key := i + j
			if key == cl {
				break
			}
			line = append(line, mat[j][key])
		}
		sort.Ints(line)
		idx := 0
		for j := 0; j < rl; j++ {
			key := i + j
			if key == cl {
				break
			}
			mat[j][key] = line[idx]
			idx++
		}
	}

	for i := 1; i < rl; i++ {
		var line []int
		for j := 0; j < cl; j++ {
			key := i + j
			if key == rl {
				break
			}
			line = append(line, mat[key][j])
		}
		sort.Ints(line)
		idx := 0
		for j := 0; j < cl; j++ {
			key := i + j
			if key == rl {
				break
			}
			mat[key][j] = line[idx]
			idx++
		}
	}
	return mat
}
