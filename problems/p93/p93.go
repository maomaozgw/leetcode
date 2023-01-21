package p93

import "strconv"

func restoreIpAddresses(s string) []string {
	var l = len(s)
	if l < 4 || l > 12 {
		return nil
	}

	result := []string{}

	result = dfs(s, 0, "", result)
	return result
}

func dfs(s string, blocks int, path string, result []string) []string {
	if blocks > 4 {
		return result
	}

	if blocks == 4 && len(s) == 0 {
		result = append(result, path[:len(path)-1])
		return result
	}
	for i := 1; i < 4; i++ {
		if i > len(s) {
			continue
		}
		num, err := strconv.Atoi(s[:i])
		if err != nil {
			return result
		}
		if string(s[:i]) == "0" || (string(s[0]) != "0" && num < 256) {
			result = dfs(s[i:], blocks+1, path+string(s[:i])+".", result)
		}
	}
	return result
}
