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

	for idx := 1; idx < n; idx++ {
		var currentArray = make([]int, 5)
		for charVal := 0; charVal < 5; charVal++ {
			for _, c := range getNext(charVal) {
				currentArray[c] = (currentArray[c] + resultArr[charVal]) % modVal
			}
		}
		copy(resultArr, currentArray)
	}
	var result int64
	for index := range resultArr {
		result = (result + int64(resultArr[index])) % int64(modVal)
	}
	return int(result)
}

func countVowelPermutationReverse(n int) int {
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

func getNext(c int) []int {
	switch c {
	case a:
		return []int{e}
	case e:
		return []int{a, i}
	case i:
		return []int{a, e, o, u}
	case o:
		return []int{i, u}
	case u:
		return []int{a}
	default:
		panic("something wrong")
	}
}
