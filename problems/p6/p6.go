package p6

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var l = len(s)
	var res = make([]byte, l)
	var it, ii = 0, 0
	var sum = 2*numRows - 2
	// first row
	for ii < l {
		res[it] = s[ii]
		it++
		ii += sum
	}
	for i := 1; i < numRows-1; i++ {
		d := sum - i*2
		ii = i
		for ii < l {
			res[it] = s[ii]
			it++
			if ii+d < l {
				res[it] = s[ii+d]
				it++
			}
			ii += sum
		}
	}
	ii = numRows - 1
	// last row
	for ii < l {
		res[it] = s[ii]
		it++
		ii += sum
	}
	return string(res)
}
