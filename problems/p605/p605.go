package p605

func canPlaceFlowers(flowerbed []int, n int) bool {
	var zc = 1
	for i := 0; i < len(flowerbed) && n > 0; i++ {
		if flowerbed[i] == 0 {
			zc++
		} else {
			zc = 0
		}
		if zc == 3 {
			n--
			zc = 1
		}
	}
	zc++
	if zc == 3 {
		n--
	}
	return n <= 0
}
