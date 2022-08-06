// https://leetcode.com/problems/poor-pigs/

package p458

import "math"

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	if buckets == 1 {
		return 0
	}

	testCount := minutesToTest/minutesToDie + 1
	for result := 1; ; result++ {
		totalCount := int(math.Pow(float64(testCount), float64(result)))
		if totalCount >= buckets {
			return result
		}
	}
	/*  我的思路
			1 只猪能从 测试次数+1 的数量的毒药中找到有问题的那瓶(如果猪没死就说明是最后一个)
	        2 只猪则能判定 (测试次数+1)*2 的毒药中找到有问题的那瓶 (行列扫描确定一个点)
	        3 只则是 (测试次数+1)*3 (x,y,z 确定一个点)
	       	推导 (测试次数+1)^猪的数量 = 最多判定的毒药个数
		看了别人的题解，这里实际上可以继续转换为：
		testCount^x >= buckets
		x >= log(buckets)/log(testCount) （看了下go 的实现 Log2 就是这么计算的 LogaN = Log(N)/Log(a)）
	*/
	// return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(testCount))))
}
