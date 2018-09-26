package mysqrt

func mySqrt(x int) int {
	if x == 1 {
		return x
	}

	var mid, start, end int
	start, end = 0, x
	for start < end {
		mid = (start + end) / 2
		if mid*mid <= x {
			if (mid+1)*(mid+1) > x {
				return mid
			}
			start = mid
		} else {
			end = mid
		}
	}
	return 0
}
