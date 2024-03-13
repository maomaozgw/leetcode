package p349

func intersection(nums1 []int, nums2 []int) []int {
	var numMap = map[int]int{}
	for _, n := range nums1 {
		numMap[n] = 0
	}
	var result []int
	for _, n := range nums2 {
		if cnt, ok := numMap[n]; ok && cnt == 0 {
			numMap[n]++
			result = append(result, n)
		}
	}
	return result
}
