package p2187

func minimumTime(time []int, totalTrips int) int64 {
	var start, end = int64(1), int64(100000000000000)
	var totalTrips64 = int64(totalTrips)
	for start <= end {
		var trip = int64(0)
		var mid = (start + end) / 2
		for i := 0; i < len(time); i++ {
			trip += mid / int64(time[i])
		}
		if trip < totalTrips64 {
			start = mid + 1
			continue
		}
		end = mid - 1
	}
	return start
}
