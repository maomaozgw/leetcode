package p2433

func findArray(pref []int) []int {
	var l = len(pref) - 1
	for ; l > 0; l-- {
		pref[l] = pref[l] ^ pref[l-1]
	}
	return pref
}
