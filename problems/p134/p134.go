package p134

func canCompleteCircuit(gas []int, cost []int) int {
	var l = len(gas)
	var totalGas, totalCost = 0, 0
	var currentGas = 0
	var result = 0
	for i := 0; i < l; i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		currentGas += gas[i] - cost[i]
		if currentGas < 0 {
			result = i + 1
			currentGas = 0
		}
	}
	if totalGas < totalCost {
		return -1
	}
	return result
}
