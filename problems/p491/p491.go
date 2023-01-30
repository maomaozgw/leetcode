package p491

import "fmt"

func findSubsequences(nums []int) [][]int {
	result, sequence := set{}, numbers{}

	backtrack(0, nums, &sequence, &result)

	return result.Value()
}

type numbers []int

func (n numbers) String() string {
	result := ""
	for _, num := range n {
		if result == "" {
			result += fmt.Sprint(num)
		} else {
			result += fmt.Sprintf(",%d", num)
		}
	}

	return result
}

func (n numbers) Value() []int {
	return append([]int{}, n...)
}

type set map[string][]int

func (s set) Value() [][]int {
	result := [][]int{}
	for _, v := range s {
		result = append(result, v)
	}

	return result
}

func backtrack(index int, nums []int, sequence *numbers, result *set) {
	if index == len(nums) {
		if len(*sequence) >= 2 {
			(*result)[(*sequence).String()] = (*sequence).Value()
		}

		return
	}
	if len(*sequence) <= 0 || (*sequence)[len(*sequence)-1] <= nums[index] {
		*sequence = append(*sequence, nums[index])
		backtrack(index+1, nums, sequence, result)
		*sequence = (*sequence)[:len(*sequence)-1]
	}
	backtrack(index+1, nums, sequence, result)
}
