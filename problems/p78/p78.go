package p78

func subsets(nums []int) [][]int {
	res := [][]int{}

	var findSubset func(idx, l int, list []int)

	findSubset = func(idx, l int, list []int) {
		if len(list) == l {
			tmp := make([]int, l)
			copy(tmp, list)
			res = append(res, tmp)
			return
		}
		for i := idx; i < len(nums); i++ {
			findSubset(i+1, l, append(list, nums[i]))
		}
	}

	for i := 0; i <= len(nums); i++ {
		findSubset(0, i, []int{})
	}
	return res

}
