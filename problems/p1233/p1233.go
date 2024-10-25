package p1233

import "sort"

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	res := []string{folder[0]}
	for i := 1; i < len(folder); i++ {
		last := res[len(res)-1]
		if len(last) > len(folder[i]) || folder[i][:len(last)] != last || folder[i][len(last)] != '/' {
			res = append(res, folder[i])
		}
	}
	return res
}
