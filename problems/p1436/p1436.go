package p1436

func destCity(paths [][]string) string {
	var nextMap = map[string]string{}
	for i := range paths {
		nextMap[paths[i][0]] = paths[i][1]
	}
	for i := range paths {
		if _, ok := nextMap[paths[i][1]]; !ok {
			return paths[i][1]
		}
	}
	return ""
}
