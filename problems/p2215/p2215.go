package p2215

func findDifference(nums1 []int, nums2 []int) [][]int {
	var set1 = make([]bool, 2001)
	var set2 = make([]bool, 2001)
	for i := range nums1 {
		set1[nums1[i]+1000] = true
	}
	for i := range nums2 {
		set2[nums2[i]+1000] = true
	}
	var result [][]int
	result = append(result, []int{})
	for i := range nums1 {
		if set2[nums1[i]+1000] {
			continue
		}
		set2[nums1[i]+1000] = true
		result[0] = append(result[0], nums1[i])

	}
	result = append(result, []int{})
	for i := range nums2 {
		if set1[nums2[i]+1000] {
			continue
		}
		set1[nums2[i]+1000] = true
		result[1] = append(result[1], nums2[i])

	}
	return result
}
