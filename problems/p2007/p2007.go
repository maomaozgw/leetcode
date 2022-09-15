package p2007

const (
	maxNum = 100000
)

func findOriginalArray(changed []int) []int {
	var l = len(changed)
	var result []int
	if l%2 != 0 {
		return result
	}
	var numCounter = make([]int, maxNum+1)
	for i := 0; i < l; i++ {
		numCounter[changed[i]]++
	}

	for i := maxNum; i >= 0; i-- {
		if numCounter[i] == 0 {
			continue
		}
		if i%2 != 0 {
			return []int{}
		}
		numCounter[i]--
		half := i / 2
		if numCounter[half] == 0 {
			return []int{}
		}
		numCounter[half]--
		result = append(result, half)
		if numCounter[i] > 0 {
			i++
		}
	}
	return result
}
