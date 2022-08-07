// https://leetcode.com/problems/count-vowels-permutation/

package p1220

const (
	a = 0
	e = 1
	i = 2
	o = 3
	u = 4
)

func countVowelPermutation(n int) int {
	var modVal = 1000000007
	var resultArr = []int{1, 1, 1, 1, 1}
	var currentArray = make([]int, 5)

	for idx := 1; idx < n; idx++ {
		currentArray[a] = (resultArr[e] + resultArr[i] + resultArr[u]) % modVal
		currentArray[e] = (resultArr[a] + resultArr[i]) % modVal
		currentArray[i] = (resultArr[e] + resultArr[o]) % modVal
		currentArray[o] = (resultArr[i]) % modVal
		currentArray[u] = (resultArr[i] + resultArr[o]) % modVal
		copy(resultArr, currentArray)
	}
	var result int64
	for index := range resultArr {
		result = (result + int64(resultArr[index])) % int64(modVal)
	}
	return int(result)
}
