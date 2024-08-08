package p885

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	var res = make([][]int, 0, rows*cols)
	res = append(res, []int{rStart, cStart})
	r, c := rStart, cStart
	step := 1
	for len(res) < rows*cols {
		for i := 0; i < step; i++ {
			c++
			if r >= 0 && r < rows && c >= 0 && c < cols {
				res = append(res, []int{r, c})
			}
		}
		for i := 0; i < step; i++ {
			r++
			if r >= 0 && r < rows && c >= 0 && c < cols {
				res = append(res, []int{r, c})
			}
		}
		step++
		for i := 0; i < step; i++ {
			c--
			if r >= 0 && r < rows && c >= 0 && c < cols {
				res = append(res, []int{r, c})
			}
		}
		for i := 0; i < step; i++ {
			r--
			if r >= 0 && r < rows && c >= 0 && c < cols {
				res = append(res, []int{r, c})
			}
		}
		step++
	}
	return res
}
